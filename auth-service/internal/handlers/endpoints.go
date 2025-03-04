package handlers

import (
	"auth-service/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HandlersManager) HandlerLogin(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (h *HandlersManager) HandlerRegister(ctx *gin.Context) {

}
