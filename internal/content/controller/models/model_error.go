/*
 * PushART - Posts | API
 *
 * API для управления постами и лентой
 */
package models

type ModelError struct {
	// Описание ошибки
	Message string `json:"message,omitempty"`
}
