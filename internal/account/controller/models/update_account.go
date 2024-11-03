package models

// Service модель изменения данных аккаунта пользователя
type UpdateAccount struct {
	// Имя пользователя
	Username string `json:"username,omitempty"`
	// Пароль пользователя
	Password string `json:"password,omitempty"`
	// Почта пользователя (если есть)
	Email string `json:"email,omitempty"`
	// Роль: читатель или автор
	Role string `json:"role"`
}
