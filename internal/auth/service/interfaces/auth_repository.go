package interfaces

import (
	"context"
	"github.com/gofrs/uuid"
)

// AuthRepository - описывает методы для работы с пользователями.
type AuthRepository interface {
	// SaveUserWithRole сохраняет пользователя в базу данных
	SaveUserWithRole(ctx context.Context, userId uuid.UUID, username string, role string, passwordHash string) error

	// UserExists проверяет, существует ли пользователь с указанным именем
	UserExists(ctx context.Context, username string) (bool, error)

	// GetUserByUserId получает пользователя по его ID
	GetUserByUserId(ctx context.Context, userId uuid.UUID) (username, role string, err error)

	// GetPasswordHashByID получает хэш пароля пользователя по его ID
	GetPasswordHashByID(ctx context.Context, userID uuid.UUID) (string, error)

	// GetUserByUsername возвращает пользователя по имени.
	GetUserByUsername(ctx context.Context, username string) (uuid.UUID, string, error)
}
