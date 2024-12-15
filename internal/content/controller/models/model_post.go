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
	PostId string `json:"postId"`
	// Заголовок поста
	Title string `json:"title"`
	// Содержимое пост
	Content string `json:"content,omitempty"`
	// Имя автора поста
	AuthorUsername string `json:"authorUsername,omitempty"`
	// Идентификатор автора поста
	AuthorId string `json:"authorId"`
	// Количество лайков поста
	Likes int `json:"likes"`
	// Поставил ли лайк текущий пользователь
	IsLiked bool `json:"isLiked"`
	// Дата создания поста
	CreatedAt string `json:"createdAt"`
	// Статус поста только для автор ми
	Status string `json:"status"`
	// Число комментариев
	NumComments int `json:"numComments"`
}

func (p *Post) String() string {
	return fmt.Sprintf("Post{PostID: %s, Title: %s, Content: %s}", p.PostId, p.Title, p.Content)
}
