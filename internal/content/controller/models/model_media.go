package models

//go:generate easyjson

//easyjson:json
type Media struct {
	// Идентификатор медиа
	MediaID string `json:"mediaID"`
	// Тип медиа
	MediaType string `json:"mediaType"`
	// URL медиа
	MediaURL string `json:"mediaURL"`
}

//easyjson:json
type MediaResponse struct {
	PostID       string   `json:"postId"`
	MediaContent []*Media `json:"mediaContent"`
}
