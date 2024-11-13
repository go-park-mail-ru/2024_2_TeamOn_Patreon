package models

// Service модель Author
type Author struct {
	Username      string
	Info          string
	Followers     int
	Subscriptions []Subscription
}
