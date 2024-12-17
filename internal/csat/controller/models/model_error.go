/*
 * PushART - СКАТ | API
 *
 * API для управления CSAT
 *
 * API version: 1.0.0
 */

package models

//go:generate easyjson -all

//easyjson:json
type ModelError struct {
	// Описание ошибки
	Message string `json:"message,omitempty"`
}
