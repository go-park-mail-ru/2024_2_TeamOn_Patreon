package models

import (
	models2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
)

// User модель репозитория
type User struct {
	UserID       models2.UserID
	Username     string
	Role         string
	PasswordHash string
}

// MapImUserToBUser конвертирует модель репозитория в модель бизнес логики
func MapImUserToBUser(user User) models2.User {
	return models2.User{
		UserID:   user.UserID,
		Username: user.Username,
		Role:     models2.Role(user.Role),
	}
}
