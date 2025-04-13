package routes

import (
	"legalmatch-api/auth"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup) {
	authGroup := rg.Group("/auth")
	{
		authGroup.POST("/login", auth.Login)
		authGroup.POST("/refresh-token", auth.RefreshToken)
		authGroup.POST("/logout", auth.Logout)

	}
}
