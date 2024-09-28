package models

const (
	AuthorStatus = "Author"
	ReaderStatus = "Reader"
)

// Profile модель репозитория
type Profile struct {
	UserID        UserID
	Username      string
	Email         string
	AvatarUrl     string
	Role          string
	Followers     int32
	Subscriptions int32
}

// UserID - ключ map`ы profiles
type UserID int
