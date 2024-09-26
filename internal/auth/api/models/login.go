package models

type Login struct {
	// Логин пользователя (имя пользователя или почта)
	Username string `json:"username,omitempty"`
	// Пароль пользователя
	Password string `json:"password"`
}
