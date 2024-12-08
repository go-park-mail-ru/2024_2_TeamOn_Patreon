package interfaces

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
)

type ContentRepository interface {
	PostInterface

	FeedInterface

	LikePostInterface

	MediaInterface

	// utils

	// GetIsLikedForPosts проставляет лайки в моделях поста
	GetIsLikedForPosts(ctx context.Context, UserID string, posts []*models.Post) error
	GetUserRole(ctx context.Context, userID string) (string, error)
	CheckCustomLayer(ctx context.Context, authorID string, layer int) (bool, error)
	// GetPostLayerBuPostID уровень поста по ид поста
	GetPostLayerByPostID(ctx context.Context, postID string) (int, error)
	// GetAuthorOfPost - получение  ID автора поста
	GetAuthorOfPost(ctx context.Context, postID string) (string, error)
	// GetUserLayerOfAuthor - уровень подписки пользователя на определенном авторе
	GetUserLayerOfAuthor(ctx context.Context, userID, authorID string) (int, error)
}

type PostInterface interface {

	// InsertPost - добавляет посты
	InsertPost(ctx context.Context, userID string, postID string, title string, content string, layer int) error
	// DeletePost - удаляет пост по id
	DeletePost(ctx context.Context, postID string) error
	// UpdateTitleOfPost - обновляет название одного поста
	UpdateTitleOfPost(ctx context.Context, postID string, title string) error
	// UpdateContentOfPost - обновляет описание одного поста
	UpdateContentOfPost(ctx context.Context, postID string, content string) error
	// GetTitleOfPost - обновляет описание одного поста
	GetTitleOfPost(ctx context.Context, postID string) (string, error)
}

type FeedInterface interface {

	// GetPopularPosts - возвращает популярные посты со смещением для анонима
	GetPopularPosts(ctx context.Context, offset int, limits int) ([]*models.Post, error)
	GetPopularPostsForUser(ctx context.Context, userID string, offset int, limits int) ([]*models.Post, error)

	// GetSubscriptionPostsForUser Подписки пользователя, сортировка по дате
	GetSubscriptionPostsForUser(ctx context.Context, userID string, offset int, limits int) ([]*models.Post, error)

	// GetAuthorPostsForLayer - подписки автора, которые может смотреть пользователь
	GetAuthorPostsForLayer(ctx context.Context, layer int, authorID string, offset, limit int) ([]*models.Post, error)
	GetAuthorPostsForAnon(ctx context.Context, authorID string, offset, limit int) ([]*models.Post, error)
	GetAuthorPostsForMe(ctx context.Context, authorID string, offset, limit int) ([]*models.Post, error)
}

type LikePostInterface interface {
	GetPostLikeID(ctx context.Context, userID string, postID string) (string, error)
	InsertLikePost(ctx context.Context, userID string, postID string) error
	DeleteLikePost(ctx context.Context, userID string, postID string) error
	GetPostLikes(ctx context.Context, postID string) (int, error)

	// SendNotificationOfLike - отправляет уведомление о лайке поста
	SendNotificationOfLike(ctx context.Context, message, userID, authorID string) error
	// GetUsername - получение имени пользователя по userID
	GetUsername(ctx context.Context, userID string) (string, error)
}

type MediaInterface interface {
	GetContentsByPost(ctx context.Context, postID string) ([]*models.Media, error)
	// SaveFile - сохраняет файл к посту
	SaveFile(ctx context.Context, postID string, file []byte, fileExtension string) error
	// DeleteFile - удаляем файл, прикреплённый к посту
	DeleteFile(ctx context.Context, postID, fileID string) error
}
