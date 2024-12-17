package models

//go:generate easyjson -all

//easyjson:json
type Subscription struct {
	AuthorID   string `json:"authorID"`
	AuthorName string `json:"authorname"`
}
