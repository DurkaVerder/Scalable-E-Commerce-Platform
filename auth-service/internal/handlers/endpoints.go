package handlers

import (
	"net/http"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/auth-service/internal/models"
	elk "github.com/DurkaVerder/Scalable-E-Commerce-Platform/auth-service/pkg/logs"

	"github.com/gin-gonic/gin"
)

func (h *HandlersManager) HandlerLogin(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.Login(user)
	if err != nil {
		elk.Log.Error("Error logging in", map[string]interface{}{
			"method": "HandlerLogin",
			"action": "Login",
			"error":  err.Error(),
		})

		if err.Error() == "not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})

}

func (h *HandlersManager) HandlerRegister(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		elk.Log.Error("Error binding JSON", map[string]interface{}{
			"method": "HandlerRegister",
			"action": "ShouldBindJSON",
			"error":  err.Error(),
		})

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Register(user); err != nil {
		elk.Log.Error("Error registering user", map[string]interface{}{
			"method": "HandlerRegister",
			"action": "Register",
			"error":  err.Error(),
		})

		if err.Error() == "invalid password" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user registered"})
}

func (h *HandlersManager) HandlerValidateToken(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")

	if err := h.service.ValidateJWT(token); err != nil {
		elk.Log.Error("Error validating token", map[string]interface{}{
			"method": "HandlerValidateToken",
			"action": "ValidateJWT",
			"error":  err.Error(),
		})

		ctx.JSON(http.StatusNonAuthoritativeInfo, gin.H{"error": "user not authorized"})
		return
	}

	userID, err := h.service.GetUserIdFromToken(token)
	if err != nil {
		elk.Log.Error("Error getting user ID from token", map[string]interface{}{
			"method": "HandlerValidateToken",
			"action": "GetUserIdFromToken",
			"error":  err.Error(),
			"token":  token,
		})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user authorized",
		"user_id": userID,
	})
}
