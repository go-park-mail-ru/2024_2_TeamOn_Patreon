/*
 * PushART - Posts | API
 *
 * API для управления постами и лентой
 */
package models

type AuthorPosts struct {
	// Идентификатор поста
	PostId string `json:"postId"`
	// Заголовок поста
	Title string `json:"title"`
	// Содержимое поста
	Content string `json:"content,omitempty"`
	// Количество лайков поста
	Likes int32 `json:"likes"`
	CreatedAt string `json:"CreatedAt"`
}
