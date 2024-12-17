/*
 * PushART - Comments | API
 *
 * API для управления комментариями к постам Порт там же где посты: 8084
 *
 * API version: 1.0.0
 */

package models

//go:generate easyjson -all

//easyjson:json
type Comment struct {
	// Уникальный идентификатор комментария
	CommentID string `json:"commentID"`
	// Содержимое коммента
	Content string `json:"content"`
	// Имя автора коммента
	Username string `json:"username"`
	// Имя пользователя кому был оставлен коммент
	UserID string `json:"userID"`
	// Время создания коммента
	CreatedAt string `json:"createdAt"`
}
