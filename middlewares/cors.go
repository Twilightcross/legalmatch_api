package middlewares

import "github.com/gin-gonic/gin"

func CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*") // 프로덕션에선 도메인 제한 권장
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		ctx.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Authorization")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	}
}
