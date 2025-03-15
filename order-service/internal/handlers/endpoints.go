package handlers

import (
	"net/http"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/order-service/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *HandlerManager) HandlerGetOrder(c *gin.Context) {
	orderID, err := h.service.ConvertStringToInt(c.Param("orderID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.service.GetOrder(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *HandlerManager) HandlerGetOrders(c *gin.Context) {
	userID, err := h.service.ConvertStringToInt(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orders, err := h.service.GetOrders(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *HandlerManager) HandlerCreateOrder(c *gin.Context) {
	userID, err := h.service.ConvertStringToInt(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var products []models.Product
	if err := c.BindJSON(&products); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.CreateOrder(userID, products)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Order created"})
}

func (h *HandlerManager) HandlerUpdateOrder(c *gin.Context) {
	orderID, err := h.service.ConvertStringToInt(c.Param("orderID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var status models.Status
	if err := c.BindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.UpdateOrder(orderID, status.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order updated"})
}
