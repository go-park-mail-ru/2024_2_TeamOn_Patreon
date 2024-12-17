/*
 * PushART - Posts | API
 *
 * API для управления постами и лентой
 */
package models

import (
	"os"
)

//go:generate easyjson -all

//easyjson:json
type MediaUploadRequest struct {
	// Айди поста, в котором нужно редактировать медиа
	PostId string `json:"postId,omitempty"`
	// Загружаемый файл (PDF, JPG, JPEG, MP3, MP4)
	File **os.File `json:"file"`
	// Формат загружаемого файла
	Format string `json:"format"`
	// Описание медиа-контента
	Description string `json:"description,omitempty"`
}
