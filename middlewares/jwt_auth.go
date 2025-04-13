package middlewares

import (
	"legalmatch-api/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid_token_format"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ParseAccessToken(tokenStr)

		if err != nil {
			if err.Error() == "expired" {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "access_token_expired",
					"message": "Access token expired. Please refresh.",
				})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid_token"})
			return
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Next()
	}
}
