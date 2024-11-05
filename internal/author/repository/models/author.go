package models

import "database/sql"

// Repository модель Author
type Author struct {
	UserID    string         `db:"user_id"`
	Username  string         `db:"username"`
	Info      sql.NullString `json:"info,omitempty"`
	Followers int
}
