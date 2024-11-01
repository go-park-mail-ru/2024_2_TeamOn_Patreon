package interfaces

import (
	"context"

	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/service/models"
)

// Интерфейс AccountRepository описывает методы взаимодействия уровня service с уровнем repository
type AccountRepository interface {
	// UserByID получает данные пользователя по указанному ID
	UserByID(ctx context.Context, userID string) (*sModels.User, error)

	// AvatarPathByID получает путь до аватара пользователя по указанному ID
	AvatarPathByID(ctx context.Context, userID string) (string, error)

	// UpdateUsername обновляет имя пользователя
	UpdateUsername(ctx context.Context, userID string, username string) error

	// UpdatePassword обновляет пароль пользователя
	UpdatePassword(ctx context.Context, userID string, hashPassword string) error

	// UpdateEmail обновляет почту пользователя
	UpdateEmail(ctx context.Context, userID string, email string) error

	// UpdateAvatar обновляет путь к аватару пользователя
	UpdateAvatar(ctx context.Context, userID string, avatarID string, avatarPath string) error

	// GenerateID генерирует ID в формате UUIDv4
	GenerateID() string
}
