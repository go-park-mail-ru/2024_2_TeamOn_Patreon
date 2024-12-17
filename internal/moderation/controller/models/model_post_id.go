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
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
)

type PostID struct {
	// ID поста
	PostID string `json:"postID"`
}

func (pID *PostID) Validate() error {
	if utils.IsValidUUIDv4(pID.PostID) {
		return nil
	}
	return global.ErrIsInvalidUUID
}

func (pID *PostID) String() string {
	return fmt.Sprintf("PostID: %s", pID.PostID)
}
