package utils

import (
	"errors"
	"fmt"
	"legalmatch-api/auth"
	"legalmatch-api/config"

	"github.com/golang-jwt/jwt/v5"
)

func ParseAccessToken(tokenStr string) (*auth.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &auth.TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return config.GetJWTSecret(), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, fmt.Errorf("expired")
		}
		return nil, err
	}

	cliams, ok := token.Claims.(*auth.TokenClaims)
	if !ok || !token.Valid {

		fmt.Println("❌ JWT invalid claims or invalid token")
		return nil, errors.New("invalid token")
	}
	return cliams, nil
}
