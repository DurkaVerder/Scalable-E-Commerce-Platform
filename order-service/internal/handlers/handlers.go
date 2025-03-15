package handlers

import (
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/order-service/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	HandlerGetOrder(c *gin.Context)
	HandlerGetOrders(c *gin.Context)
	HandlerCreateOrder(c *gin.Context)
	HandlerUpdateOrder(c *gin.Context)
}

type HandlerManager struct {
	service service.Service
}

func NewHandlerManager(service service.Service) *HandlerManager {
	return &HandlerManager{
		service: service,
	}
}
