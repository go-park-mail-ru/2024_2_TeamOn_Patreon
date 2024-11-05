package models

import (
	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/service/models"
	valid "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/validate"
)

// Controller модель страницы автора
type AuthorPage struct {
	// Имя пользователя
	Username string `json:"authorUsername"`
	// Информация о себе
	Info string `json:"info,omitempty"`
	// Количество подписчиков
	Followers int `json:"followers"`
	// Подписки автора
	Subscriptions []sModels.Subscription `json:"subscriptions"`
}

// MapAuthorToAuthorPage конвертирует модель автора в модель controller страницы автора
func MapAuthorToAuthorPage(author sModels.Author, subscriptions []sModels.Subscription) AuthorPage {
	return AuthorPage{
		Username:      valid.Sanitize(author.Username),
		Info:          valid.Sanitize(author.Info),
		Followers:     author.Followers,
		Subscriptions: subscriptions,
	}
}
