package interfaces

import (
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/buisness/models"
)

// UserRepository - описывает методы для работы с пользователями.
type UserRepository interface {
	// SaveUser сохраняет пользователя в базу данных
	SaveUser(username string, role int, passwordHash string) (*bModels.User, error)

	// UserExists проверяет, существует ли пользователь с указанным именем
	UserExists(username string) (bool, error)

	// GetUserByID получает пользователя по его ID
	GetUserByID(userID int) (*bModels.User, error)

	// GetPasswordHashByID получает хэш пароля пользователя по его ID
	GetPasswordHashByID(userID int) (string, error)

	// GetUserByUsername возвращает пользователя по имени.
	GetUserByUsername(username string) (*bModels.User, error)
}
