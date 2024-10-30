package interfaces

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	"github.com/gofrs/uuid"
)

type ContentRepository interface {
	InsertPost(userId uuid.UUID, postId uuid.UUID, title string, content string, layer int) error
	// GetPopularPosts - возвращает популярные посты со смещением для анонима
	GetPopularPosts(offset int, limits int) ([]models.Post, error)
	GetPopularPostsForUser(userId uuid.UUID, offset int, limits int) ([]models.Post, error)

	UpdatePost(authorId uuid.UUID, postID uuid.UUID, post models.Post) error
	GetAuthorByPost(postID uuid.UUID) (uuid.UUID, error)

	// LikePost

	IsLikePutPost(userId uuid.UUID, postID uuid.UUID) (bool, error)
	InsertLikePost(userId uuid.UUID, postID uuid.UUID) error
	DeleteLikePost(userId uuid.UUID, postID uuid.UUID) error
	GetPostLikes(postID uuid.UUID) (int, error)
}
