package models

type Token struct {
	// JWT-токен для авторизованного пользователя
	Token string `json:"token,omitempty"`
}
