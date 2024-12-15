// pkg

package models

import "time"

type Post struct {
	// Идентификатор поста
	PostID string
	// Заголовок поста
	Title string
	// Содержимое поста
	Content string
	// Имя автора поста
	AuthorUsername string
	// Идентификатор автора поста
	AuthorID string
	// Количество лайков поста
	Likes int
	// Уровень подписки
	Layer int
	// Лайкнул ли текущий пользователь
	IsLiked bool
	// Когда создан
	CreatedDate time.Time
	// Непустой только для выборки автором для авторов
	Status string
	// Число комментариев
	NumComments int
}
