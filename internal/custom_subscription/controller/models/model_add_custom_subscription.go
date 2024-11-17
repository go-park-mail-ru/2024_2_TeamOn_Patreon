package models

// AddCustomSubscription Данные для добавления кастомной подписки, которые отправляет фронт
type AddCustomSubscription struct {
	// Заголовок кастномной (платной) подписки
	Title string `json:"title"`
	// Описание кастомной (платной) подписки
	Description string `json:"description,omitempty"`
	// Стоимость  кастомной (платной) подписки в рублях
	Cost int32 `json:"cost"`
	// Уровень кастомной (платной) подписки. Уровень [0:3]
	Layer int32 `json:"layer"`
}
