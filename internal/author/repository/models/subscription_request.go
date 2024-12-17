package models

type SubscriptionRequest struct {
	SubReqID   string
	UserID     string
	AuthorID   string
	MonthCount int
	Layer      int
}
