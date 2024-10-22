/*
 * PushART - Posts | API
 *
 * API для управления постами и лентой
 *
 */
package models

type Post struct {
	// Идентификатор поста
	PostId string `json:"postId,omitempty"`
	// Заголовок поста
	Title string `json:"title,omitempty"`
	// Содержимое поста
	Content string `json:"content,omitempty"`
	// Имя автора поста
	AuthorUsername string `json:"authorUsername,omitempty"`
	// Идентификатор автора поста
	AuthorId string `json:"authorId,omitempty"`
	// Количество лайков поста
	Likes int32 `json:"likes,omitempty"`
}
