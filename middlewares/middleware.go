package middlewares

import "github.com/gin-gonic/gin"

func Apply(r *gin.Engine) {
	r.Use(
		Logger(),
		CORS(),
	)
}
