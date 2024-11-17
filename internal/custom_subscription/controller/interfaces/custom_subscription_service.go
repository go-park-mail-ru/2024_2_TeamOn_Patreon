package interfaces

import "context"

type CustomSubscriptionService interface {
	CreateCustomSub(ctx context.Context, userID string, title, description string, layer, cost int) error
}
