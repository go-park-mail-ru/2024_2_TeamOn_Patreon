package models

import "time"

type Comment struct {
	CommentID string
	Content   string
	Username  string
	UserID    string
	CreatedAt time.Time
}
