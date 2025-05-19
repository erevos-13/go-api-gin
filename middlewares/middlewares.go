package middlewares

import (
	"net/http"

	"example.com/gin-api/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userId, err := utils.ValidateJwtToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return

	}
	context.Set("userId", userId)
	context.Next()
}
