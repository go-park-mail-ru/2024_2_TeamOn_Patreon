package interfaces

import "context"

type ModerationService interface {

	// ComplaintPost Жалоба на пост
	ComplaintPost(ctx context.Context, postID string, userID string) error

	// DecisionPost - Решение от модератора
	DecisionPost(ctx context.Context, postID string, userID string) error

	// GetPostsForModeration
	// TODO: Добавить бизнес модель поста для отображения модератору!!! в возвращаемых значениях
	GetPostsForModeration(ctx context.Context, userID string, filter string, limit, offset int) error
}
