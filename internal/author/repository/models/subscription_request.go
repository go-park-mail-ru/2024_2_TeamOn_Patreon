package models

type SubscriptionRequest struct {
	UserID     string
	AuthorID   string
	MonthCount int
	Layer      int
}
