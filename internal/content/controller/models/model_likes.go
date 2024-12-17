/*
 * PushART - Posts | API
 *
 * API для управления постами и лентой
 */
package models

//go:generate easyjson

//easyjson:json
type Likes struct {
	// Количество лайков
	Count int `json:"count"`
}
