package models

//go:generate easyjson -all

// ModelError Сообщение об ошибке. Возвыращает бэк
//
//easyjson:json
type ModelError struct {
	// Описание ошибки
	Message string `json:"message,omitempty"`
}
