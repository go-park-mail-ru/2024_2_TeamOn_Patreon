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
	Followers     uint
	Subscriptions uint
	PostsAmount   uint
	PostTitle     string
	PostContent   string
	PostDate      string
}

// UserID - ключ map`ы profiles
type UserID int
