package interfaces

import (
	bmAccount "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/repository/models"
)

// AccountRepository - описывает методы для работы с аккаунтом.
type AccountRepository interface {
	// SaveAccount сохраняет профиль в базу данных
	SaveAccount(userID int, username string, role int) (*bmAccount.Account, error)

	// AccountExists проверяет, существует ли пользователь с указанным ID
	AccountExists(userID int) (bool, error)

	// GetAccountByID получает профиль по ID пользователя
	GetAccountByID(userID int) (*bmAccount.Account, error)
}
