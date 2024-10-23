package models

const (
	AuthorStatus = "Author"
	ReaderStatus = "Reader"
)

// Repository модель Account
type Account struct {
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

// UserID - ключ map`ы Accounts
type UserID int
