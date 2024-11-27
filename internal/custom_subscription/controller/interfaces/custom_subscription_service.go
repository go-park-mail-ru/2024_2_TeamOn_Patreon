package interfaces

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/pkg/models"
)

type CustomSubscriptionService interface {
	CustomSubscription
	SearchAuthor
}

type CustomSubscription interface {
	CreateCustomSub(ctx context.Context, userID string, title, description string, layer, cost int) error

	GetLayerForNewCustomSub(ctx context.Context, userID string) ([]*models.SubscriptionLayer, error)
	GetCustomSubscription(ctx context.Context, authorID string) ([]*models.CustomSubscription, error)
}

type SearchAuthor interface {
	SearchAuthor(ctx context.Context, searchTerm string, limit, offset int) ([]string, error)
}
