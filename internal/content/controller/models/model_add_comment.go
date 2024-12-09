/*
 * PushART - Comments | API
 *
 * API для управления комментариями к постам Порт там же где посты: 8084
 *
 * API version: 1.0.0
 */

package models

type AddComment struct {
	// Содержимое коммента
	CommentID string `json:"commentID"`
}
