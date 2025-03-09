package handlers

import (
	"net/http"
	"strconv"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/cart-service/internal/models"
	elk "github.com/DurkaVerder/Scalable-E-Commerce-Platform/cart-service/pkg/logs"

	"github.com/gin-gonic/gin"
)

// HandlerAddProduct adds a product to the cart.
func (h *HandlersManager) HandlerAddProduct(c *gin.Context) {
	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	userID, err := h.service.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user"})
		return
	}

	if err := h.service.AddProductToCart(userID, product.ID, product.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product added to cart"})
}

// HandlerGetCart returns the cart of a user.
func (h *HandlersManager) HandlerGetCart(c *gin.Context) {
	userID, err := h.service.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user"})
		return
	}

	products, err := h.service.GetCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// HandlerDeleteProduct removes a product from the cart.
func (h *HandlersManager) HandlerDeleteProduct(c *gin.Context) {
	userID, err := h.service.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user"})
		return
	}

	productID, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		elk.Log.Error("Failed to convert the product ID to an integer: "+err.Error(), map[string]interface{}{
			"method": "HandlerDeleteProduct",
			"action": "converting the product ID to an integer",
		})
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	if err := h.service.RemoveProductFromCart(userID, productID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product removed from cart"})
}

// HandlerUpdateProduct updates the quantity of a product in the cart.
func (h *HandlersManager) HandlerUpdateProduct(c *gin.Context) {
	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	userID, err := h.service.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user"})
		return
	}

	if err := h.service.UpdateProductQuantity(userID, product.ID, product.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product updated in cart"})
}
