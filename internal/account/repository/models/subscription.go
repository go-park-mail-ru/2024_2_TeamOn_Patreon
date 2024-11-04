package models

type Subscription struct {
	AuthorID   string `db:"user_id"`
	AuthorName string `db:"username"`
}
