package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

// TokenClaims
// Структура, которая содержит поля для JWT токена
type TokenClaims struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
