package handlers

import (
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/catalog-service/internal/service"
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	HandlerAllProducts(ctx *gin.Context)
	HandlerProductByID(ctx *gin.Context)
}

type HandlersManager struct {
	service service.Service
}

func NewHandlersManager(service service.Service) *HandlersManager {
	return &HandlersManager{
		service: service,
	}
}
