package models

type Media struct {
	// Идентификатор медиа
	MediaID string `json:"mediaID"`
	// Тип медиа
	MediaType string `json:"mediaType"`
	// URL медиа
	MediaURL string `json:"mediaURL"`
}

type MediaResponse struct {
	PostId       string   `json:"postId"`
	MediaContent []*Media `json:"mediaContent"`
}
