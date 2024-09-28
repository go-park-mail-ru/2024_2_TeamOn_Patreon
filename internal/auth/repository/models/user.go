package models

import (
	models2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/buisness/models"
)

// User модель репозитория
type User struct {
	UserID       UserID
	Username     string
	Role         int
	PasswordHash string
}

// UserID - ключ мапы users
type UserID int

// MapImUserToBUser конвертирует модель репозитория в модель бизнес логики
func MapImUserToBUser(user User) models2.User {
	return models2.User{
		UserID:   int(user.UserID),
		Username: user.Username,
		Role:     models2.Role(user.Role),
	}
}
