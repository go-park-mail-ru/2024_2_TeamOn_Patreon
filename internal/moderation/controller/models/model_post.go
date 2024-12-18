/*
 * PushART - Модерация | API
 *
 * API для интерфейса модератора по проверке постов  ПОРТ  8087
 *
 * API version: 1.0.1
 */

package models

import "fmt"

//go:generate easyjson

//easyjson:json
type Post struct {
	// Идентификатор поста
	PostID string `json:"postID"`
	// Заголовок поста
	Title string `json:"title"`
	// Содержимое поста
	Content string `json:"content,omitempty"`
	// Имя автора поста
	AuthorUsername string `json:"authorUsername"`
	// Идентификатор автора поста
	AuthorID string `json:"authorID"`
	// Статус поста
	Status string `json:"status"`
	// Время создания поста
	CreatedAt string `json:"createdAt"`
}

func (p *Post) String() string {
	return fmt.Sprintf("Post{Post ID: %s\nTitle: %s\nStatus: %s\n AuthorUsername:%s}",
		p.PostID, p.Title, p.Status, p.AuthorUsername)
}
