package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		method := ctx.Request.Method
		path := ctx.Request.URL.Path

		ctx.Next()

		duration := time.Since(start)
		status := ctx.Writer.Status()

		fmt.Printf("[LOGGER] %s %s → %d (%v)\n", method, path, status, duration)
	}
}
