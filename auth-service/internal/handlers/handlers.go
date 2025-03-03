package handlers

import (
	"auth-service/internal/service"

	"github.com/gin-gonic/gin"
)

type Handlers interface {
	HandlerLogin(ctx *gin.Context)
	HandlerRegister(ctx *gin.Context)
}

type HandlersManager struct {
	service service.Service
}

func NewHandlersManager(service service.Service) *HandlersManager {
	return &HandlersManager{service: service}
}
