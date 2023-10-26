package handlers

import (
	"github.com/devminnu/tek-job-portal/internal/services"
	"github.com/devminnu/tek-job-portal/pkg/auth"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	UserHandler
}

type UserHandler interface {
	Register(*gin.Context)
	Login(*gin.Context)
}

type CompanyHandler interface{}

type handler struct {
	UserHandler
	svc  services.Service
	auth *auth.Auth
}

func New(a *auth.Auth, s services.Service) Handler {
	return &handler{
		auth: a,
		svc:  s,
	}
}
