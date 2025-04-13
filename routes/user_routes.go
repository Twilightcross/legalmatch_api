package routes

import (
	"legalmatch-api/controllers"
	"legalmatch-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", controllers.CreateUser)
		userGroup.GET("/myinfo", middlewares.JwtAuth(), controllers.GetMyInfo)
	}
}
