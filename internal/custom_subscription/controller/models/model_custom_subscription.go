package models

import "fmt"

//go:generate easyjson

// CustomSubscription Данные о кастономной подписки, которые возвращает бэк
//
//easyjson:json
type CustomSubscription struct {
	// Идентификатор кастомной (платной) подписки
	CustomSubscriptionID string `json:"customSubscriptionID"`
	// Заголовок кастномной (платной) подписки
	Title string `json:"title"`
	// Описание кастомной (платной) подписки
	Description string `json:"description,omitempty"`
	// Стоимость  кастомной (платной) подписки в рублях в месяц
	Cost int `json:"cost"`
	// Уровень кастомной (платной) подписки
	Layer int `json:"layer"`
}

func (sc CustomSubscription) String() string {
	return fmt.Sprintf("Custom Subscription ID: %v, Title: %s, Description: %s, Cost: %d, Layer: %d",
		sc.CustomSubscriptionID, sc.Title, sc.Description, sc.Cost, sc.Layer)
}
