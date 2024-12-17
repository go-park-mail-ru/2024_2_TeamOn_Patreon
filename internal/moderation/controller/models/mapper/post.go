package mapper

import (
	models2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/moderation/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/moderation/pkg/models"
)

func MapBPostToTPost(post *models.Post) *models2.Post {
	return &models2.Post{
		PostID:         post.PostID,
		Title:          post.Title,
		Content:        post.Content,
		AuthorUsername: post.AuthorUsername,
		AuthorID:       post.AuthorID,
		Status:         post.Status,
		CreatedAt:      post.CreatedAt.String(),
	}
}

func MapBPostsToTPosts(posts []*models.Post) []*models2.Post {
	result := make([]*models2.Post, 0, len(posts))
	for _, post := range posts {
		post := MapBPostToTPost(post)
		result = append(result, post)
	}
	return result
}
