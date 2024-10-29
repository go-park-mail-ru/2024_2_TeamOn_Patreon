package imagine

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	"github.com/gofrs/uuid"
)

type Post struct {
	postId    uuid.UUID
	title     string
	content   string
	layer     int
	contentId uuid.UUID
	authorID  uuid.UUID
	likes     int
}

func MapRepositoryPostToPkgPost(post Post) models.Post {
	return models.Post{
		PostId:   post.postId.String(),
		Title:    post.title,
		Content:  post.content,
		Likes:    post.likes,
		AuthorId: post.authorID.String(),
	}
}
