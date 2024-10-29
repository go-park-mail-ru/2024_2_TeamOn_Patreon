package interfaces

import (
	"context"

	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/service/models"
	b2Models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
)

// Интерфейс AccountRepository описывает методы взаимодействия уровня service с уровнем repository
type AccountRepository interface {
	// AccountExists проверяет, существует ли пользователь с указанным ID
	AccountExist(ctx context.Context, userID string) (bool, error)

	// SaveAccount сохраняет аккаунт в базу данных
	// Этот метод скорее всего реализован в сервисе AUTH, но это не точно
	SaveAccount(ctx context.Context, userID string, username string, role b2Models.Role) (*bModels.Account, error)

	// GetAccountByID получает профиль по ID пользователя
	FindByID(ctx context.Context, userID string) (*bModels.Account, error)
}
