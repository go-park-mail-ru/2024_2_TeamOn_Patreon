/*
 * PushART - Posts | API
 *
 * API для управления постами и лентой
 */
package models

import "fmt"

//go:generate easyjson

//easyjson:json
type ModelError struct {
	// Описание ошибки
	Message string `json:"message,omitempty"`
}

func (me *ModelError) String() string {
	return fmt.Sprintf("ModelError {Message: %s}", me.Message)
}
