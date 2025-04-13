package auth

import (
	"errors"
	"legalmatch-api/config"
	"legalmatch-api/models"
	"time"
)

func GenerateRefreshToken(userID uint) (string, error) {
	rt := models.RefreshToken{
		UserID: userID,
	}

	if err := config.DB.Create(&rt).Error; err != nil {
		return "", err
	}
	return rt.Token, nil
}

func ValidateRefreshToken(token string) (*models.RefreshToken, error) {
	var rt models.RefreshToken
	if err := config.DB.Where("token = ?", token).First(&rt).Error; err != nil {
		return nil, err
	}

	if rt.RevokedAt != nil || rt.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("refresh token is invalid or expired")
	}

	return &rt, nil
}

func RevokeRefreshToken(token string) error {
	var rt models.RefreshToken
	if err := config.DB.Where("token = ?", token).First(&rt).Error; err != nil {
		return err
	}

	now := time.Now()
	rt.RevokedAt = &now
	return config.DB.Save(&rt).Error
}
