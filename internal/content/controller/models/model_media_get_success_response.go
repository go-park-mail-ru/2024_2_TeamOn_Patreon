/*
 * PushART - Posts | API
 *
 * API для управления постами и лентой
 */
package models

//go:generate easyjson -all

//easyjson:json
type MediaGetSuccessResponse struct {
	PostId string `json:"postId,omitempty"`

	MediaContent []MediaGetSuccessResponseMediaContent `json:"mediaContent,omitempty"`
}
