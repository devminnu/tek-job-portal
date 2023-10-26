package routes

// import (
// 	"fmt"

// 	"github.com/rs/zerolog/log"

// 	"time"

// 	"github.com/devminnu/tek-job-portal/internal/handlers"
// 	"github.com/devminnu/tek-job-portal/internal/middlewares"
// 	"github.com/devminnu/tek-job-portal/internal/services"
// 	"github.com/gin-gonic/gin"

// 	"net/http"
// )

// // Define a function called API that takes an argument a of type *auth.Auth
// // and returns a pointer to a gin.Engine

// func API(r *gin.Engine) (router *gin.Engine) {

// 	// Attempt to create new middleware with authentication
// 	// Here, *auth.Auth passed as a parameter will be used to set up the middleware
// 	m, err := middlewares.NewMiddleware(a)
// 	// If there is an error in setting up the middleware, panic and stop the application
// 	// then log the error message
// 	if err != nil {
// 		log.Panic().Msg("middlewares not set up")

// 		return
// 	}
// 	s := services.New(c)
// 	h := handlers.New(a, s)

// 	// If there is an error in setting up the middleware, panic and stop the application
// 	// then log the error message
// 	// if err != nil {
// 	// 	log.Panic().Msg("middlewares not set up")
// 	// }

// 	// Attach middleware's Log function and Gin's Recovery middleware to our application
// 	// The Recovery middleware recovers from any panics and writes a 500 HTTP response if there was one.
// 	router.Use(m.Log(), gin.Recovery())

// 	// Define a route at path "/check"
// 	// If it receives a GET request, it will use the m.Authenticate(check) function.
// 	// r.GET("/check", m.Authenticate(check))
// 	// r.POST("/signup", h.Signup)
// 	// r.POST("/login", h.Login)
// 	// r.POST("/add", m.Authenticate(h.AddInventory))
// 	// r.POST("/view", m.Authenticate(h.ViewInventory))

// 	router.POST("/api/user/register", h.Register)
// 	router.POST("/api/user/login", h.Login)

// 	// Return the prepared Gin engine
// 	return
// }

// func check(c *gin.Context) {
// 	//handle panic using recovery function when happening in separate goroutine
// 	//go func() {
// 	//	panic("some kind of panic")
// 	//}()
// 	time.Sleep(time.Second * 3)
// 	select {
// 	case <-c.Request.Context().Done():
// 		fmt.Println("user not there")
// 		return
// 	default:
// 		c.JSON(http.StatusOK, gin.H{"msg": "statusOk"})

// 	}

// }
