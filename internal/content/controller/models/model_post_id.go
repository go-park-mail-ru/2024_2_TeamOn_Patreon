package models

import "fmt"

//go:generate easyjson

//easyjson:json
type PostId struct {
	// id post
	PostId string `json:"postId"`
}

func (p PostId) String() string {
	return fmt.Sprintf("PostID:\t %s", p.PostId)
}
