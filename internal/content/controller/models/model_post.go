/*
 * PushART - Posts | API
 *
 * API для управления постами и лентой
 *
 */
package models

import "fmt"

type Post struct {
	// Идентификатор поста
	PostId string `json:"post_id"`
	// Заголовок поста
	Title string `json:"title"`
	// Содержимое пост
	Content string `json:"content,omitempty"`
	// Имя автора поста
	AuthorUsername string `json:"author_username,omitempty"`
	// Идентификатор автора поста
	AuthorId string `json:"author_id"`
	// Количество лайков поста
	Likes int `json:"likes"`
	// Уровень подписки
	Layer int `json:"layer"`
	// Поставил ли лайк текущий пользователь
	IsLiked bool `json:"is_liked"`
}

func (p *Post) String() string {
	return fmt.Sprintf("Post{PostId: %s, Title: %s, Content: %s}", p.PostId, p.Title, p.Content)
}
