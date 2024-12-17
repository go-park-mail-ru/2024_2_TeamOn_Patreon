/*
 * PushART - Модерация | API
 *
 * API для интерфейса модератора по проверке постов  ПОРТ  8087
 *
 * API version: 1.0.1
 */

package models

//go:generate easyjson -all

//easyjson:json
type ModelError struct {
	// Описание ошибки
	Message string `json:"message,omitempty"`
}
