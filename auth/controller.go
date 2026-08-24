package auth

import (
	"legalmatch-api/config"
	"legalmatch-api/models"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "요청 형식이 잘못되었습니다"})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "이메일 또는 비밀번호가 일치하지 않습니다"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "이메일 또는 비밀번호가 일치하지 않습니다"})
		return
	}

	token, err := GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		log.Println("Access Token 생성 실패 이유:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Access Token 생성 실패"})
		return
	}

	refreshToken, err := GenerateRefreshToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Refresh Token 생성 실패"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token, "refresh_token": refreshToken})
}

func RefreshToken(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")

	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_token_format"})
		return
	}

	refreshToken := strings.TrimPrefix(authHeader, "Bearer ")

	var rt models.RefreshToken
	err := config.DB.Where("token = ?", refreshToken).First(&rt).Error
	if err != nil || rt.IsRevoked() || rt.IsExpired() {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid_or_expired_refresh_token"})
		return
	}
	var user models.User
	if err := config.DB.Where("id = ?", rt.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user_not_found"})
		return
	}

	accessToken, err := GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed_to_generate_access_token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}

func Logout(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "refresh_token이 필요합니다"})
		return
	}

	if err := RevokeRefreshToken(req.RefreshToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "유효하지 않은 refresh_token 입니다"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "로그아웃 완료"})
}
