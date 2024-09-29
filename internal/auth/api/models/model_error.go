package models

import "fmt"

type ModelError struct {
	// Сообщение об ошибке
	Message string `json:"message,omitempty"`
}

func (me *ModelError) String() string {
	return fmt.Sprintf("ModelError {Message: %s}", me.Message)
}
