package models

// Service модель аккаунта пользователя
type Account struct {
	// Имя пользователя
	Username string `json:"username"`
	// Почта пользователя (если есть)
	Email string `json:"email,omitempty"`
	// Роль: читатель или автор
	Role string `json:"role"`
	// Подписки пользователя
	Subscriptions []Subscription `json:"subscriptions"`
}
