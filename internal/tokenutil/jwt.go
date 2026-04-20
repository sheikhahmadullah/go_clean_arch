package tokenutil

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// this will generate tokens for us
var secret_key = []byte("secret-key")

func GenerateToken(userId uint) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret_key)
}
