package handlers

import (
	"cart-service/internal/service"

	"github.com/gin-gonic/gin"
)

type Handlers interface {
	HandlerAddProduct(ctx *gin.Context)
	HandlerGetCart(ctx *gin.Context)
	HandlerDeleteProduct(ctx *gin.Context)
	HandlerUpdateProduct(ctx *gin.Context)
}

type HandlersManager struct {
	service service.Service
}

func NewHandlersManager(service service.Service) *HandlersManager {
	return &HandlersManager{
		service: service,
	}
}
