package handlers

import (
	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/payment-service/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	HandlerCreatePayment(c *gin.Context)
}

type HandlerManager struct {
	service service.Service
}

func NewHandlerManager(service service.Service) *HandlerManager {
	return &HandlerManager{
		service: service,
	}
}
