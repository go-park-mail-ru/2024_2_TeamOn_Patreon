package models

import "time"

type Post struct {
	PostID         string
	Title          string
	Content        string
	AuthorID       string
	AuthorUsername string
	Status         string
	CreatedAt      time.Time
}
