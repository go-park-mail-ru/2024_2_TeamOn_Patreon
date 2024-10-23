package models

const (
	AuthorStatus = "Author"
	ReaderStatus = "Reader"
)

// API Модель аккаунта пользователя
type Account struct {
	// Имя пользователя
	Username string `json:"username"`
	// Почта пользователя (если есть)
	Email string `json:"email,omitempty"`
	// Роль: читатель или автор
	Role string `json:"role"`
	// Количество подписок
	Subscriptions uint `json:"subscriptions"`
}
