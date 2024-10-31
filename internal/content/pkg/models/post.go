// pkg

package models

import "time"

type Post struct {
	// Идентификатор поста
	PostId string
	// Заголовок поста
	Title string
	// Содержимое поста
	Content string
	// Имя автора поста
	AuthorUsername string
	// Идентификатор автора поста
	AuthorId string
	// Количество лайков поста
	Likes int
	// Уровень подписки
	Layer int
	// Лайкнул ли текущий пользователь
	IsLiked bool
	// Когда создан
	CreatedDate time.Time
}
