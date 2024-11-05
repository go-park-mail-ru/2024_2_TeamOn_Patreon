/*
 * PushART - Posts | API
 *
 * API для управления постами и лентой
 */
package models

import (
	"os"
)

type MediaGetSuccessResponseMediaContent struct {
	MediaId string `json:"mediaId,omitempty"`
	// Формат файла
	MediaType string `json:"mediaType,omitempty"`

	Description string `json:"description,omitempty"`
	// Загрузка медиа-файла
	File **os.File `json:"file,omitempty"`
}
