package models

import "database/sql"

// Service модель автора
type Author struct {
	Username      string
	Info          sql.NullString // так как может быть пустым
	Followers     int
	Subscriptions []Subscription
}
