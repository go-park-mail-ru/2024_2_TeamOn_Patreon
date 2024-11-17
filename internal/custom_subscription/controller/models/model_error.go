package models

// ModelError Сообщение об ошибке. Возвыращает бэк
type ModelError struct {
	// Описание ошибки
	Message string `json:"message,omitempty"`
}
