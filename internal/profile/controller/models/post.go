package models

import (
	"time"
)

// Модель поста
type Post struct {
	// ID поста
	Id uint32 `json:"id"`
	// Текст поста
	Text string `json:"text"`
	// Ссылка на медиа-контент поста
	MediaContentUrl string `json:"media_content_url,omitempty"`
	// Дата создания поста
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// Дата обновления поста
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
