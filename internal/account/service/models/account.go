package bmodels

const (
	AuthorStatus = "Author"
	ReaderStatus = "Reader"
)

// Repository модель Account
type Account struct {
	UserID        UserID
	Username      string
	Email         string
	Role          string
	Subscriptions []Subscription
}

// UserID - ключ map`ы Accounts
type UserID string
