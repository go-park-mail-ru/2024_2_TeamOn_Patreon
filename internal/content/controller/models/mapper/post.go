package mapper

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/models"
	models2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/validate"
)

func MapInterfacePostToTransportPost(post models2.Post) *models.Post {
	title := validate.Sanitize(post.Title)
	content := validate.Sanitize(post.Content)
	authorUsername := validate.Sanitize(post.AuthorUsername)

	return &models.Post{
		PostId:         post.PostID,
		Title:          title,
		Content:        content,
		AuthorUsername: authorUsername,
		AuthorId:       post.AuthorID,
		Likes:          post.Likes,
		IsLiked:        post.IsLiked,
		CreatedAt:      post.CreatedDate.String(),
	}
}

func MapCommonPostsToControllerPosts(posts []*models2.Post) []*models.Post {
	tPosts := make([]*models.Post, 0, len(posts))
	for _, post := range posts {
		tPosts = append(tPosts, MapInterfacePostToTransportPost(*post))
	}
	return tPosts
}
