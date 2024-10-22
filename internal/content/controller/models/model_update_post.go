/*
 * PushART - Posts | API
 *
 * API для управления постами и лентой
 */
package models

// Все, чего нет - не заменяется
type UpdatePost struct {
	// Идентификатор поста для обновления
	PostId int32 `json:"postId"`
	// Заголовок поста
	Title string `json:"title,omitempty"`
	// Содержимое поста (текст)
	Content string `json:"content,omitempty"`
	// Уровень на котором можно смотреть пост
	Layer int32 `json:"layer,omitempty"`
}
