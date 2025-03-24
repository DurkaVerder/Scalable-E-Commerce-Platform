package handlers

import (
	"net/http"

	elk "github.com/DurkaVerder/elk-send-logs/elk"
	"github.com/gin-gonic/gin"
)

func (h *HandlersManager) HandlerAllProducts(ctx *gin.Context) {
	category := ctx.Query("category")

	products, err := h.service.GetProducts(category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	elk.Log.SendMsg(
		elk.LogMessage{
			Level:   'I',
			Message: "Products fetched",
			Fields: map[string]interface{}{
				"method":   "HandlerAllProducts",
				"action":   "GetProducts",
				"category": category,
			},
		})

	ctx.JSON(http.StatusOK, products)
}

func (h *HandlersManager) HandlerProductByID(ctx *gin.Context) {
	id := ctx.Param("id")

	product, err := h.service.GetProductByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	elk.Log.SendMsg(
		elk.LogMessage{
			Level:   'I',
			Message: "Product fetched",
			Fields: map[string]interface{}{
				"method": "HandlerProductByID",
				"action": "GetProductByID",
				"id":     id,
			},
		})

	ctx.JSON(http.StatusOK, product)
}
