package interfaces

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	"github.com/gofrs/uuid"
)

type ContentRepository interface {
	// post

	// InsertPost - добавляет посты
	InsertPost(ctx context.Context, userId uuid.UUID, postId uuid.UUID, title string, content string, layer int) error
	// DeletePost - удаляет пост по id
	DeletePost(ctx context.Context, postID uuid.UUID) error

	// GetAuthorOfPost - получение  ID автора поста
	GetAuthorOfPost(ctx context.Context, postID uuid.UUID) (uuid.UUID, error)

	UpdateTitleOfPost(ctx context.Context, postID uuid.UUID, title string) error
	UpdateContentOfPost(ctx context.Context, postID uuid.UUID, content string) error

	// UpdatePost(authorId uuid.UUID, postID uuid.UUID, post models.Post) error

	// subscription

	CheckCustomLayer(ctx context.Context, authorId uuid.UUID, layer int) (bool, error)

	// user

	GetUserRole(ctx context.Context, userId uuid.UUID) (string, error)

	// GetPopularPosts - возвращает популярные посты со смещением для анонима
	GetPopularPosts(offset int, limits int) ([]models.Post, error)
	GetPopularPostsForUser(userId uuid.UUID, offset int, limits int) ([]models.Post, error)

	GetAuthorByPost(postID uuid.UUID) (uuid.UUID, error)

	// LikePost

	IsLikePutPost(userId uuid.UUID, postID uuid.UUID) (bool, error)
	InsertLikePost(userId uuid.UUID, postID uuid.UUID) error
	DeleteLikePost(userId uuid.UUID, postID uuid.UUID) error
	GetPostLikes(postID uuid.UUID) (int, error)
}
