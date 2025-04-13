package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader("Authorization")
		fmt.Println("DEBUG token:", authHeader)
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization 헤더가 필요합니다"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := ValidateToken(tokenStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "토큰이 유효하지 않거나 만료되었습니다"})
			return
		}

		//역할 확인
		allowed := false
		for _, role := range allowedRoles {
			if claims.Role == role {
				allowed = true
				break
			}
		}
		if !allowed {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "접근 권한이 없습니다"})
			return
		}

		// 유저 정보 context에 저장
		ctx.Set("user_id", claims.UserID)
		ctx.Set("role", claims.Role)
		ctx.Next()
	}

}
