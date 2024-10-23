package jwt

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/golang-jwt/jwt/v5"
)

// TokenClaims
// Структура, которая содержит поля для JWT токена
type TokenClaims struct {
	UserID   string      `json:"userID"`
	Username string      `json:"username"`
	Role     models.Role `json:"role"`
	jwt.RegisteredClaims
}
