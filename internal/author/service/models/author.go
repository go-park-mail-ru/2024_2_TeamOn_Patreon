package models

// Author модель автора
type Author struct {
	Username      string
	Info          string
	Followers     int
	Subscriptions []Subscription
}
