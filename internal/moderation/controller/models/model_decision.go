/*
 * PushART - Модерация | API
 *
 * API для интерфейса модератора по проверке постов  ПОРТ  8087
 *
 * API version: 1.0.1
 */

package models

// Decision решение модератора об одобрении поста
type Decision struct {
	// Идентификатор поста
	PostID string `json:"postID"`
	// Новый статус поста
	Status string `json:"status"`
}
