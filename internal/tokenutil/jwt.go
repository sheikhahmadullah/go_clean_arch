package tokenutil

import (
	"golang-rest-api/bootstrap"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// this will generate tokens for us

func GenerateToken(userId uint) (string, error) {
	secret := []byte(bootstrap.GetEnv("JWT_SECRET"))
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}
