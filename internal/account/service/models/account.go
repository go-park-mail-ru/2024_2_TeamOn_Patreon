package models

import "database/sql"

const (
	AuthorStatus = "Author"
	ReaderStatus = "Reader"
)

// Service модель User
type User struct {
	UserID        string
	Username      string
	Email         sql.NullString // так как может быть пустым
	Role          string
	Subscriptions []Subscription // ?? оставить здесь или вынести отдельно ??
}
