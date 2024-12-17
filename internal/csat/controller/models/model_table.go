/*
 * PushART - Posts | API
 *
 * API для управления постами и лентой
 *
 * API version: 1.0.0
 */

package models

//go:generate easyjson

//easyjson:json
type StatTable struct {
	// Тема вопорса
	Theme string `json:"theme"`
	// Рейтинг средний
	Rating string `json:"rating"`
}
