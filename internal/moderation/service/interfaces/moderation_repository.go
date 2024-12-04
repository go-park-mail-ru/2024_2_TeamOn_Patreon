package interfaces

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/moderation/pkg/models"
)

type ModerationRepository interface {
	User
	Status
}

type User interface {
	// GetPostLayerBuPostID уровень поста по ид поста
	GetPostLayerBuPostID(ctx context.Context, postID string) (int, error)
	// GetAuthorOfPost - получение  ID автора поста
	GetAuthorOfPost(ctx context.Context, postID string) (string, error)
	// GetUserLayerOfAuthor - уровень подписки пользователя на определенном авторе
	GetUserLayerOfAuthor(ctx context.Context, userID, authorID string) (int, error)

	GetUserRole(ctx context.Context, userID string) (string, error)
}

type Status interface {
	UpdatePostStatus(ctx context.Context, postID string, status string) error

	GetPostsByStatus(ctx context.Context, status string, limit, offset int) ([]*models.Post, error)

	GetStatusByPostID(ctx context.Context, postID string) (string, error)
}
