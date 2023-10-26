package main

import (
	"context"
	"crypto/rsa"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/devminnu/tek-job-portal/internal/database"
	"github.com/devminnu/tek-job-portal/internal/handlers"
	"github.com/devminnu/tek-job-portal/internal/middlewares"
	"github.com/devminnu/tek-job-portal/internal/services"
	"github.com/devminnu/tek-job-portal/pkg/auth"
	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("starting app")
	err := startApp()
	if err != nil {
		log.Panic().Err(err).Send()
	}
	log.Info().Msg("hello this is our app")
}

func startApp() error {
	// Initialize authentication support
	log.Info().Msg("main : Started : Initializing authentication support")
	privateKey, publicKey, err := getRSAKeys()
	if err != nil {
		panic(err)
	}
	a, err := auth.NewAuth(privateKey, publicKey)
	if err != nil {
		return fmt.Errorf("constructing auth %w", err)
	}

	// Start Database
	log.Info().Msg("main : Started : Initializing db support")
	db, err := database.Open()
	if err != nil {
		return fmt.Errorf("connecting to db %w", err)
	}
	pg, err := db.DB()
	if err != nil {
		return fmt.Errorf("Failed to get database instance: %w ", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = pg.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("Database is not connected: %w ", err)
	}
	dbConn, err := database.New(db)
	if err != nil {
		return fmt.Errorf("connecting to db %w", err)
	}
	log.Info().Msg("db connection successful")

	svc, err := services.New(dbConn)
	if err != nil {
		return err
	}
	err = svc.AutoMigrate()
	if err != nil {
		return err
	}
	// Create a new Gin engine; Gin is a HTTP web framework written in Go
	router := gin.New()
	h := handlers.New(a, svc)
	registerMiddlewares(router)
	registerRoutes(router, h)

	// Initialize http service
	api := http.Server{
		Addr:         ":8080",
		ReadTimeout:  8000 * time.Second,
		WriteTimeout: 800 * time.Second,
		IdleTimeout:  800 * time.Second,
		Handler:      router,
	}

	// channel to store any errors while setting up the service
	serverErrors := make(chan error, 1)
	go func() {
		log.Info().Str("port", api.Addr).Msg("main: API listening")
		serverErrors <- api.ListenAndServe()
	}()
	//shutdown channel intercepts ctrl+c signals
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error %w", err)
	case sig := <-shutdown:
		log.Info().Msgf("main: Start shutdown %s", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		//Shutdown gracefully shuts down the server without interrupting any active connections.
		//Shutdown works by first closing all open listeners, then closing all idle connections,
		//and then waiting indefinitely for connections to return to idle and then shut down.
		err := api.Shutdown(ctx)
		if err != nil {
			//Close immediately closes all active net.Listeners
			err = api.Close() // forcing shutdown
			return fmt.Errorf("could not stop server gracefully %w", err)
		}
	}

	return nil
}

func registerMiddlewares(router *gin.Engine) {
	m := middlewares.NewLoggerMiddleware()
	router.Use(m.Log(), gin.Recovery())
}

func registerRoutes(router *gin.Engine, h handlers.Handler) {
	// user routes
	router.POST("/api/user/register", h.Register)
	router.POST("/api/user/login", h.Login)
}

func getRSAKeys() (privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey, err error) {
	privatePEM, err := os.ReadFile("private.pem")
	if err != nil {
		err = fmt.Errorf("reading auth private key %w", err)

		return
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
	if err != nil {
		err = fmt.Errorf("parsing auth private key %w", err)

		return
	}

	publicPEM, err := os.ReadFile("pubkey.pem")
	if err != nil {
		err = fmt.Errorf("reading auth public key %w", err)

		return
	}

	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicPEM)
	if err != nil {
		err = fmt.Errorf("parsing auth public key %w", err)

		return
	}

	return
}
