package jwt

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func TestCreateToken_Basic(t *testing.T) {

	secret := "super-secret-key-123456"
	service := NewJWTService(secret)
	userID := "user12345678"
	token := service.CreateToken(userID)

	// 2. Тестовые данные - разные случаи userID
	testCases := []struct {
		token     string
		name      string
		userID    string
		setupTest func(t *testing.T, token string)
	}{
		{

			name:   "success token",
			userID: userID,
			token:  token,
			setupTest: func(t *testing.T, token string) {
				parsedToken, _ := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
					return []byte(secret), nil
				})
				// токен валиден
				require.Equal(t, parsedToken.Valid, true)

				tk, _, _ := jwt.NewParser().ParseUnverified(token, jwt.MapClaims{})
				claims := tk.Claims.(jwt.MapClaims)
				ParsedUserID := claims["user_id"].(string)
				//ID юзеров совпадают
				require.Equal(t, userID, ParsedUserID)

			},
		},
		{
			name:   "empty userID",
			userID: "",
			setupTest: func(t *testing.T, token string) {
				parser := new(jwt.Parser)
				tokenObj, _, err := parser.ParseUnverified(token, jwt.MapClaims{})
				require.NoError(t, err)

				// Проверяем что в токене пустой user_id
				if claims, ok := tokenObj.Claims.(jwt.MapClaims); ok {
					require.Equal(t, "", claims["user_id"])
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			token := service.CreateToken(tc.userID)

			if tc.setupTest != nil {
				tc.setupTest(t, token)
			}
		})
	}
}
