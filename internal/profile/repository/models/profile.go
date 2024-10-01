package models

const (
	AuthorStatus = "Author"
	ReaderStatus = "Reader"
)

// Repository модель Profile
type Profile struct {
	UserID        UserID
	Username      string
	Email         string
	Status        string
	AvatarUrl     string
	Role          string
	Followers     int
	Subscriptions int
	PostsAmount   int
}

// UserID - ключ map`ы profiles
type UserID int
