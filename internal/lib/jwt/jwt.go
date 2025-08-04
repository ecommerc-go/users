package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// CreateJWTToken создание токена с подписью
func CreateJWTToken(id string, secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // Срок - 24 часа
	})

	// Подписываем секретом
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}
	return tokenString
}
