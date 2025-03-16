package handlers

import (
	"net/http"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/payment-service/internal/models"
	elk "github.com/DurkaVerder/Scalable-E-Commerce-Platform/payment-service/pkg/logs"
	"github.com/gin-gonic/gin"
)

func (h *HandlerManager) HandlerCreatePayment(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		elk.Log.Error("Error binding JSON", map[string]interface{}{
			"method": "HandlerCreatePayment",
			"action": "binding JSON",
			"order":  order,
			"error":  err.Error(),
		})
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pi, err := h.service.CreatePaymentIntent(order)
	if err != nil {
		elk.Log.Error("Error creating payment", map[string]interface{}{
			"method": "HandlerCreatePayment",
			"action": "creating payment",
			"order":  order,
			"error":  err.Error(),
		})
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	elk.Log.Info("Payment created", map[string]interface{}{
		"method": "HandlerCreatePayment",
		"action": "creating payment",
		"order":  order,
	})

	c.JSON(http.StatusCreated, gin.H{"client_secret": pi.ClientSecret})
}
