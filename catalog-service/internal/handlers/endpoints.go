package handlers

import (
	"net/http"

	elk "github.com/DurkaVerder/Scalable-E-Commerce-Platform/catalog-service/pkg/logs"
	"github.com/gin-gonic/gin"
)

func (h *HandlersManager) HandlerAllProducts(ctx *gin.Context) {
	category := ctx.Query("category")

	products, err := h.service.GetProducts(category)
	if err != nil {
		elk.Log.Error("Error getting products", map[string]interface{}{
			"method":   "HandlerAllProducts",
			"action":   "GetProducts",
			"category": category,
			"error":    err.Error(),
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	elk.Log.Info("Products fetched", map[string]interface{}{
		"method":   "HandlerAllProducts",
		"action":   "GetProducts",
		"category": category,
		"count":    len(products),
	})

	ctx.JSON(http.StatusOK, products)
}

func (h *HandlersManager) HandlerProductByID(ctx *gin.Context) {
	id := ctx.Param("id")

	product, err := h.service.GetProductByID(id)
	if err != nil {
		elk.Log.Error("Error getting product by ID", map[string]interface{}{
			"method":    "HandlerProductByID",
			"action":    "GetProductByID",
			"productId": id,
			"error":     err.Error(),
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	elk.Log.Info("Product fetched by ID", map[string]interface{}{
		"method":    "HandlerProductByID",
		"action":    "GetProductByID",
		"productId": id,
	})

	ctx.JSON(http.StatusOK, product)
}
