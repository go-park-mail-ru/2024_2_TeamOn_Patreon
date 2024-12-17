package interfaces

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/moderation/pkg/models"
)

type ModerationService interface {

	// ComplaintPost Жалоба на пост
	ComplaintPost(ctx context.Context, postID string, userID string) error

	// DecisionPost - Решение от модератора
	DecisionPost(ctx context.Context, postID string, userID string, status string) error

	// GetPostsForModeration - возвращает посты необходимые для модерации по фильтру
	GetPostsForModeration(ctx context.Context, userID string, filter string, limit, offset int) ([]*models.Post, error)
}
