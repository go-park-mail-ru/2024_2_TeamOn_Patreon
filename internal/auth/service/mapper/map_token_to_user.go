package mapper

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service/jwt"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
)

// MapTokenToUser
// Функция для маппинга
func MapTokenToUser(token *jwt.TokenClaims) bModels.User {
	return bModels.User{
		UserID:   bModels.UserID(token.UserID),
		Username: token.Username,
		Role:     token.Role,
	}
}
