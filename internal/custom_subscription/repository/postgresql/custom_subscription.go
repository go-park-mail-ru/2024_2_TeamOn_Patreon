package postgresql

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
)

const (
	// getExistLayersForAuthorSQL достает уровни, на которых у автора есть подписки
	// Input: $1 - authorID
	// Output: []layer (int)
	getExistLayersForAuthorSQL = `
select layer, default_layer_name
from Custom_Subscription
JOIN Subscription_Layer USING (subscription_layer_id)
where author_id=$1
`

	// getFreeLayersForAuthorSQL достает уровни, на которых у автора нет подписки
	// Input: $1 - authorID
	// Output: []layer (int)
	getFreeLayersForAuthorSQL = `
select layer, default_layer_name
from Subscription_Layer
where subscription_layer_id not in (
	select subscription_layer_id 
	from Custom_Subscription 
	where author_id=$1
)
 
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

	// getCustomSubscriptionSQL - возвращает кастомные подписки автора
	// Input: $1 - authorID
	// Output: CumSubID, title, description, cost, layer
	getCustomSubscriptionSQl = `
SELECT custom_subscription_id, custom_name, info, cost, layer 
	FROM public.custom_subscription
	join subscription_layer using (subscription_layer_id)
	where author_id=$1
`
)

// GetLayersForAuthor - возвращает все уровни подписки, на которых
// у автора уже создана подписка!
func (csr *CustomSubscriptionRepository) GetLayersForAuthor(ctx context.Context, authorID string) ([]*models.SubscriptionLayer, error) {
	op := "internal.custom_subscription.repository.postgresql.GetLayersForAuthor"

	rows, err := csr.db.Query(ctx, getExistLayersForAuthorSQL, authorID)
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

// GetFreeLayersForAuthor - возвращает все уровни подписки, на которых
// у автора нет подписки!!!
func (csr *CustomSubscriptionRepository) GetFreeLayersForAuthor(ctx context.Context, authorID string) ([]*models.SubscriptionLayer, error) {
	op := "internal.custom_subscription.repository.postgresql.GetLayersForAuthor"

	rows, err := csr.db.Query(ctx, getFreeLayersForAuthorSQL, authorID)
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
	op := "internal.custom_subscription.repository.postgresql.CreateCustomSub"

	_, err := csr.db.Exec(ctx, insertCustomSubscriptionSQL, userID, title, cost, description, layer)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

func (csr *CustomSubscriptionRepository) GetCustomSubscription(ctx context.Context, authorID string) ([]*models.CustomSubscription, error) {
	op := "internal.custom_subscription.repository.postgresql.GetCustomSubscription"

	rows, err := csr.db.Query(ctx, getCustomSubscriptionSQl, authorID)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	defer rows.Close()

	customSubs := make([]*models.CustomSubscription, 0)

	var (
		cumSubID    string
		title       string
		description string
		cost        int
		layer       int
	)

	for rows.Next() {
		if err = rows.Scan(&cumSubID, &title, &description, &cost, &layer); err != nil {
			return nil, errors.Wrap(err, op)
		}
		customSubs = append(customSubs, &models.CustomSubscription{
			CustomSubscriptionID: cumSubID,
			Title:                title,
			Description:          description,
			Cost:                 cost,
			Layer:                layer,
		})
	}
	return customSubs, nil
}
