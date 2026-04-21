package middleware

import (
	"golang-rest-api/bootstrap"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := []byte(bootstrap.GetEnv("JWT_SECRET"))
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no token"})
			c.Abort()
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// ✅ extract claims
		claims := token.Claims.(jwt.MapClaims)

		userID := uint(claims["user_id"].(float64))

		// ✅ store in context
		c.Set("user_id", userID)

		c.Next()

	}
}
