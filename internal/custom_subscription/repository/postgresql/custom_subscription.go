package postgresql

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
)

const (
	// getLayersForAuthorSQL достает уровни, на которых у автора есть подписки
	// Input: $1 - authorID
	// Output: []layer (int)
	getLayersForAuthorSQL = `
select layer, default_layer_name
from Custom_Subscription
JOIN Subscription_Layer USING (subscription_layer_id)
where author_id=$1
`

	// getUserRoleSQL достает роль пользователя
	// Input: $1 - userID
	// Output: role (string)
	getUserRoleSQL = `
select role_default_name
from Role
join People USING(role_id)
where user_id = $1
`

	// insertCustomSubscriptionSQL - вставляет запись с кастомной подпиской
	// Input: $1 - authorID, $2 - title, $3 - cost, $4 - description, $5 - layer (int)
	insertCustomSubscriptionSQL = `
INSERT INTO Custom_Subscription (custom_subscription_id, author_id, custom_name, cost, info, subscription_layer_id, created_date) VALUES
    (gen_random_uuid(), $1, $2, $3, $4, (SELECT subscription_layer_id FROM Subscription_Layer WHERE layer = $5), NOW())
 
`
)

// GetLayersForAuthor - возвращает все уровни подписки, на которых
// у автора уже создана подписка!
func (csr *CustomSubscriptionRepository) GetLayersForAuthor(ctx context.Context, authorID string) ([]*models.SubscriptionLayer, error) {
	op := "internal.custom_subscription.repository.postgresql.GetLayersForAuthor"

	rows, err := csr.db.Query(ctx, getLayersForAuthorSQL, authorID)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	defer rows.Close()

	layers := make([]*models.SubscriptionLayer, 0)

	var (
		layer     int
		layerName string
	)

	for rows.Next() {
		if err = rows.Scan(&layer, &layerName); err != nil {
			return nil, errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op, "Got sub layer layer=%v name=%v for author=%v", layer, layerName, authorID)
		layers = append(layers, &models.SubscriptionLayer{Layer: layer, LayerName: layerName})
	}

	return layers, nil
}

func (csr *CustomSubscriptionRepository) GetUserRole(ctx context.Context, userID string) (string, error) {
	op := "internal.custom_subscription.repository.postgresql.GetUserRole"

	rows, err := csr.db.Query(ctx, getUserRoleSQL, userID)
	if err != nil {
		return "", errors.Wrap(err, op)
	}
	defer rows.Close()

	var (
		role string
	)

	for rows.Next() {
		if err = rows.Scan(&role); err != nil {
			return "", errors.Wrap(err, op)
		}
		return role, nil
	}

	return "", nil
}

func (csr *CustomSubscriptionRepository) CreateCustomSub(ctx context.Context, userID, title, description string, layer, cost int) error {
	return nil
}
