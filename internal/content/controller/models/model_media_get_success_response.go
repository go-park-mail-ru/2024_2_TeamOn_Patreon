/*
 * PushART - Posts | API
 *
 * API для управления постами и лентой
 */
package models

type MediaGetSuccessResponse struct {
	PostId string `json:"postId,omitempty"`

	MediaContent []MediaGetSuccessResponseMediaContent `json:"mediaContent,omitempty"`
}
