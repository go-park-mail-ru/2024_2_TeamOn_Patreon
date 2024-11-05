package models

// Service модель автора
type Author struct {
	Username      string
	Info          string
	Followers     int
	Subscriptions []Subscription
}
