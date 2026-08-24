package utils

import (
	"errors"
	"fmt"
	"legalmatch-api/auth"
	"legalmatch-api/config"

	"github.com/golang-jwt/jwt/v5"
)

func ParseAccessToken(tokenStr string) (*auth.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &auth.JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return config.GetJWTSecret(), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, fmt.Errorf("expired")
		}
		return nil, err
	}

	claims, ok := token.Claims.(*auth.JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
