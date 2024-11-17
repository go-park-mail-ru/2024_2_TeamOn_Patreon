package models

// CustomSubscription Данные о кастономной подписки, которые возвращает бэк
type CustomSubscription struct {
	// Идентификатор кастомной (платной) подписки
	CustomSubscriptionID string `json:"customSubscriptionID,omitempty"`
	// Заголовок кастномной (платной) подписки
	Title string `json:"title,omitempty"`
	// Описание кастомной (платной) подписки
	Description string `json:"description,omitempty"`
	// Стоимость  кастомной (платной) подписки в рублях в месяц
	Cost int `json:"cost,omitempty"`
	// Уровень кастомной (платной) подписки
	Layer int `json:"layer,omitempty"`
}
