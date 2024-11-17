package postgresql

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/pkg/models"
)

// GetLayersForAuthor - возвращает все уровни подписки, на которых
// у автора уже создана подписка!
func (csr *CustomSubscriptionRepository) GetLayersForAuthor(ctx context.Context, authorID string) ([]*models.SubscriptionLayer, error) {
	return nil, nil
}

func (csr *CustomSubscriptionRepository) GetUserRole(ctx context.Context, userID string) (string, error) {
	return "", nil
}

func (csr *CustomSubscriptionRepository) CreateCustomSub(ctx context.Context, userID, title, description string, layer, cost int) error {
	return nil
}
