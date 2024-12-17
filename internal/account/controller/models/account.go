package models

import (
	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/service/models"
	valid "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/validate"
)

//go:generate easyjson

// Service модель аккаунта пользователя
//
//easyjson:json
type Account struct {
	// Имя пользователя
	Username string `json:"username"`
	// Почта пользователя (если есть)
	Email string `json:"email,omitempty"`
	// Роль: читатель или автор
	Role string `json:"role"`
	// Подписки пользователя
	Subscriptions []sModels.Subscription `json:"subscriptions"`
}

// MapUserToAccount конвертирует модель пользователя с подписками в модель controller
func MapUserToAccount(user sModels.User, subscriptions []sModels.Subscription) Account {
	return Account{
		Username:      valid.Sanitize(user.Username),
		Email:         valid.Sanitize(user.Email),
		Role:          valid.Sanitize(user.Role),
		Subscriptions: subscriptions,
	}
}
