package routes

import "github.com/gin-gonic/gin"

func RegisterAllRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		RegisterUserRoutes(v1)
		RegisterAuthRoutes(v1)
	}
}
