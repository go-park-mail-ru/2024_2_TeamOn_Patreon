package mapper

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/polina-auth/internal/auth/behavior/jwt"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/polina-auth/internal/common/buisness/models"
)

// MapTokenToUser
// Функция для маппинга
func MapTokenToUser(token *jwt.TokenClaims) bModels.User {
	return bModels.User{
		UserID:   token.UserID,
		Username: token.Username,
		Role:     token.Role,
	}
}
