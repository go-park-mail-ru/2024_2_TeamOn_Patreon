package mapper

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/models"
	models2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
)

func MapInterfacePostToTransportPost(post models2.Post) *models.Post {
	return &models.Post{
		PostId:         post.PostId,
		Title:          post.Title,
		Content:        post.Content,
		AuthorUsername: post.AuthorUsername,
		AuthorId:       post.AuthorId,
		Likes:          post.Likes,
		IsLiked:        post.IsLiked,
	}
}

func MapCommonPostsToControllerPosts(posts []models2.Post) []models.Post {
	tPosts := make([]models.Post, 0, len(posts))
	for _, post := range posts {
		tPosts = append(tPosts, *MapInterfacePostToTransportPost(post))
	}
	return tPosts
}
