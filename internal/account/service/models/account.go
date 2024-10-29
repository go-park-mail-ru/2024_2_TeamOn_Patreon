package models

const (
	AuthorStatus = "Author"
	ReaderStatus = "Reader"
)

// Service модель User
type User struct {
	UserID        string
	Username      string
	Email         string
	Role          string
	Subscriptions []Subscription // ?? оставить здесь или вынести отдельно ??
}
