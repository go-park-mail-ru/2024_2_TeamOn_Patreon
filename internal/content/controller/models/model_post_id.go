package models

import "fmt"

type PostId struct {
	// id post
	PostId string `json:"postId"`
}

func (p PostId) String() string {
	return fmt.Sprintf("PostId:\t %s", p.PostId)
}
