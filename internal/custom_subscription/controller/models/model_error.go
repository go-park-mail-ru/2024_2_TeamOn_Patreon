package models

//go:generate easyjson

// ModelError Сообщение об ошибке. Возвыращает бэк
//
//easyjson:json
type ModelError struct {
	// Описание ошибки
	Message string `json:"message,omitempty"`
}
