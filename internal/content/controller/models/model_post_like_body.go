/*
 * PushART - Posts | API
 *
 * API для управления постами и лентой
 */
package models

type PostLikeBody struct {
	// ID поста, который нужно лайкнуть или убрать лайк
	PostId int32 `json:"postId"`
}
