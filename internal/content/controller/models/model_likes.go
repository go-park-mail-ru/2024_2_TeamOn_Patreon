/*
 * PushART - Posts | API
 *
 * API для управления постами и лентой
 */
package models

//go:generate easyjson -all

//easyjson:json
type Likes struct {
	// Количество лайков
	Count int `json:"count"`
}
