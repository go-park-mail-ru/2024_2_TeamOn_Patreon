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
	Followers     int32
	Subscriptions int32
	PostsAmount   int32
}

// UserID - ключ map`ы profiles
type UserID int
