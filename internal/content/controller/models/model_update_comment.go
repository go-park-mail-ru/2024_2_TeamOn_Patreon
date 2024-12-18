/*
 * PushART - Comments | API
 *
 * API для управления комментариями к постам Порт там же где посты: 8084
 *
 * API version: 1.0.0
 */

package models

import sanitize "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/validate"

//go:generate easyjson

//easyjson:json
type UpdateComment struct {
	// Содержимое коммента
	Content string `json:"content"`
}

func (uc *UpdateComment) Validate() (bool, error) {
	uc.Content = sanitize.Sanitize(uc.Content)
	if len(uc.Content) > 500 {
		uc.Content = string([]rune(uc.Content)[:500])
	}
	return true, nil
}
