package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secret string
}

func NewJWTService(secret string) *JWTService {
	return &JWTService{secret: secret}
}

// CreateToken создает JWT токен для пользователя
func (s *JWTService) CreateToken(userID string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	// Подписываем секретом
	tokenString, err := token.SignedString([]byte(s.secret))
	if err != nil {
		panic(err)
	}
	return tokenString
}
