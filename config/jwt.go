package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env 파일이 없습니다. 운영 환경에서는 무시해도 됩니다.")
	}
}

func GetJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("❌ JWT_SECRET 환경변수가 설정되지 않았습니다.")

	}
	return []byte(secret)
}
