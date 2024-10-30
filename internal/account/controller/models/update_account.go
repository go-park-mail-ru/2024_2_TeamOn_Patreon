package models

// Service модель изменения данных аккаунта пользователя
type UpdateAccount struct {
	// Имя пользователя
	Username string `json:"username"`
	// Пароль пользователя
	Password string `json:"password"`
	// Почта пользователя (если есть)
	Email string `json:"email,omitempty"`
	// Роль: читатель или автор
	Role string `json:"role"`
}
