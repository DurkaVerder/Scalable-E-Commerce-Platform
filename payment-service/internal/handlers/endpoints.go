package handlers

import (
	"net/http"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/payment-service/internal/models"
	elk "github.com/DurkaVerder/elk-send-logs/elk"
	"github.com/gin-gonic/gin"
)

func (h *HandlerManager) HandlerCreatePayment(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		elk.Log.SendMsg(elk.LogMessage{
			Level:   'E',
			Message: "Failed to bind JSON",
			Fields: map[string]interface{}{
				"method": "HandlerCreatePayment",
				"action": "bind_json",
				"error":  err.Error(),
			},
		})

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pi, err := h.service.CreatePaymentIntent(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	elk.Log.SendMsg(elk.LogMessage{
		Level:   'I',
		Message: "Payment intent created",
		Fields: map[string]interface{}{
			"method":   "HandlerCreatePayment",
			"action":   "create_payment_intent",
			"amount":   order.Amount,
			"user_id":  order.UserID,
			"order_id": order.ID,
		},
	})

	c.JSON(http.StatusCreated, gin.H{"client_secret": pi.ClientSecret})
}
