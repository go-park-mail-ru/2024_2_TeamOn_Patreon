package models

import (
	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/service/models"
)

// Service модель аккаунта пользователя
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
		Username:      user.Username,
		Email:         user.Email,
		Role:          user.Role,
		Subscriptions: subscriptions,
	}
}
