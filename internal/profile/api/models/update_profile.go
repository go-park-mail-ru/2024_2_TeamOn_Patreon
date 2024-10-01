package models

type UpdateProfile struct {
	// Новое имя пользователя
	Username string `json:"username,omitempty"`
	// Новая почта пользователя
	Email string `json:"email,omitempty"`
}
