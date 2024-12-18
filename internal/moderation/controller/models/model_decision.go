/*
 * PushART - Модерация | API
 *
 * API для интерфейса модератора по проверке постов  ПОРТ  8087
 *
 * API version: 1.0.1
 */

package models

import (
	"fmt"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/moderation/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
)

//go:generate easyjson

// Decision решение модератора об одобрении поста
//
//easyjson:json
type Decision struct {
	// Идентификатор поста
	PostID string `json:"postID"`
	// Новый статус поста
	Status string `json:"status"`
}

func (d *Decision) Validate() error {
	if !d.validatePostID() {
		return global.ErrIsInvalidUUID
	}
	if !d.validateStatus() {
		return global.ErrStatusIncorrect
	}
	return nil
}

func (d *Decision) String() (result string) {
	return fmt.Sprintf("Decision{PostID: %s\tStatus: %s}", d.PostID, d.Status)
}

func (d *Decision) validatePostID() bool {
	return utils.IsValidUUIDv4(d.PostID)
}

func (d *Decision) validateStatus() bool {
	return models.CheckStatus(d.Status)
}
