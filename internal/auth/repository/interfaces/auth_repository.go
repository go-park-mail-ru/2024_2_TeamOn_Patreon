package interfaces

import (
	bSession "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/behavior/session"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/business/models"
)

// AuthRepository - описывает методы для работы с пользователями.
type AuthRepository interface {
	UserRepository
	SessionRepository
}

type UserRepository interface {
	// SaveUser сохраняет пользователя в базу данных
	SaveUser(username string, role int, passwordHash string) (*bModels.User, error)

	// UserExists проверяет, существует ли пользователь с указанным именем
	UserExists(username string) (bool, error)

	// GetUserByID получает пользователя по его ID
	GetUserByID(userID string) (*bModels.User, error)

	// GetPasswordHashByID получает хэш пароля пользователя по его ID
	GetPasswordHashByID(userID string) (string, error)

	// GetUserByUsername возвращает пользователя по имени.
	GetUserByUsername(username string) (*bModels.User, error)

	// RemoveUserByID удаляет пользователя по ID
	RemoveUserByID(userID string) error
}

type SessionRepository interface {
	SaveSession(session bSession.SessionModel) (string, error)

	DeleteSession(session bSession.SessionModel) error

	CheckSession(session bSession.SessionModel) (bool, error)
}
