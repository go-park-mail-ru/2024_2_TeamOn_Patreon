/*
 * PushART - Posts | API
 *
 * API для управления постами и лентой
 */
package models

type AddPost struct {
	// Заголовок поста
	Title string `json:"title"`
	// Содержимое поста
	Content string `json:"content"`
	// Уровень подписки, на котором можно смотреть пост, по умолчанию - для всех
	Layer int32 `json:"layer,omitempty"`
}
