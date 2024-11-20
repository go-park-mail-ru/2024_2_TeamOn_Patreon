package interfaces

import "context"

type CustomSubscriptionService interface {
	CustomSubscription
	SearchAuthor
}

type CustomSubscription interface {
	CreateCustomSub(ctx context.Context, userID string, title, description string, layer, cost int) error
}

type SearchAuthor interface {
	SearchAuthor(ctx context.Context, searchTerm string, limit, offset int) ([]string, error)
}
