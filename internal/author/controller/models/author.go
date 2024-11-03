package models

// Controller модель автора
type Author struct {
	// Имя пользователя
	Username string `json:"authorUsername"`
	// Информация о себе
	Info string `json:"info,omitempty"`
	// Количество подписчиков
	Followers int `json:"followers"`
	// Подписки автора
	Subscriptions []Subscription `json:"subscriptions"`
}
