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
	// UpdateTitleOfPost - обновляет название одного поста
	UpdateTitleOfPost(ctx context.Context, postID uuid.UUID, title string) error
	// UpdateContentOfPost - обновляет описание одного поста
	UpdateContentOfPost(ctx context.Context, postID uuid.UUID, content string) error

	// feed

	// GetPopularPosts - возвращает популярные посты со смещением для анонима
	GetPopularPosts(ctx context.Context, offset int, limits int) ([]*models.Post, error)
	GetPopularPostsForUser(ctx context.Context, userId uuid.UUID, offset int, limits int) ([]*models.Post, error)

	// GetSubscriptionPostsForUser Подписки пользователя, сортировка по дате
	GetSubscriptionPostsForUser(ctx context.Context, userId uuid.UUID, offset int, limits int) ([]*models.Post, error)

	// GetAuthorPostsForUser - подписки автора, которые может смотреть пользователь
	GetAuthorPostsForLayer(ctx context.Context, layer int, authorId uuid.UUID, offset, limit int) ([]*models.Post, error)
	GetAuthorPostsForAnon(ctx context.Context, authorId uuid.UUID, offset, limit int) ([]*models.Post, error)
	GetAuthorPostsForMe(ctx context.Context, authorId uuid.UUID, offset, limit int) ([]*models.Post, error)

	// GetIsLikedForPosts проставляет лайки в моделях поста
	GetIsLikedForPosts(ctx context.Context, UserId uuid.UUID, posts []*models.Post) error

	// LikePost

	GetPostLikeId(ctx context.Context, userId uuid.UUID, postID uuid.UUID) (uuid.UUID, error)
	InsertLikePost(ctx context.Context, userId uuid.UUID, postID uuid.UUID) error
	DeleteLikePost(ctx context.Context, userId uuid.UUID, postID uuid.UUID) error
	GetPostLikes(ctx context.Context, postID uuid.UUID) (int, error)

	// utils

	GetUserRole(ctx context.Context, userId uuid.UUID) (string, error)
	CheckCustomLayer(ctx context.Context, authorId uuid.UUID, layer int) (bool, error)
	GetAuthorByPost(postID uuid.UUID) (uuid.UUID, error)
	// GetPostLayerBuPostId уровень поста по ид поста
	GetPostLayerBuPostId(ctx context.Context, postID uuid.UUID) (int, error)
	// GetAuthorOfPost - получение  ID автора поста
	GetAuthorOfPost(ctx context.Context, postID uuid.UUID) (uuid.UUID, error)
	// GetUserLayerOfAuthor - уровень подписки пользователя на определенном авторе
	GetUserLayerOfAuthor(ctx context.Context, userId, authorId uuid.UUID) (int, error)
}
