package interfaces

import (
	"context"

	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/service/models"
)

// Интерфейс AccountRepository описывает методы взаимодействия уровня service с уровнем repository
type AccountRepository interface {
	// UserByID получает данные пользователя по указанному ID
	UserByID(ctx context.Context, userID string) (*sModels.User, error)
}
