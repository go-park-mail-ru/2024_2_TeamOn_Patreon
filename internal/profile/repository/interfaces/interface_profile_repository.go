package interfaces

import (
	bmProfile "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/profile/repository/models"
)

// ProfileRepository - описывает методы для работы с профилем.
type UserRepository interface {
	// SaveProfile сохраняет профиль в базу данных
	SaveProfile(userID int, username string, role int) (*bmProfile.Profile, error)

	// UserExists проверяет, существует ли пользователь с указанным ID
	UserExists(userID int) (bool, error)

	// GetProfileByID получает профиль по ID пользователя
	GetProfileByID(userID int) (*bmProfile.Profile, error)
}
