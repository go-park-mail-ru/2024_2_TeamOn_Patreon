package postgresql

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/pkg/models"
	"github.com/pkg/errors"
)

const (
	// getCustomSubscriptionSQL - возвращает кастомные подписки автора
	// Input: $1 - authorID, $2 - title
	// Output: CumSubID, title, description, cost, layer
	getCustomSubscriptionByTitleSQl = `
SELECT custom_subscription_id, custom_name, info, cost, layer 
	FROM public.custom_subscription
	join subscription_layer using (subscription_layer_id)
	where author_id=$1 and custom_name=$2
`
)

func (csr *CustomSubscriptionRepository) GetCustomSubscriptionByTitle(ctx context.Context, authorID string, title string) (*models.CustomSubscription, error) {
	op := "custom_subscription.repository.getCustomSubscriptionByTitle"

	rows, err := csr.db.Query(ctx, getCustomSubscriptionByTitleSQl, authorID, title)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	defer rows.Close()

	var (
		cumSubID    string
		description string
		cost        int
		layer       int
	)

	for rows.Next() {
		if err = rows.Scan(&cumSubID, &title, &description, &cost, &layer); err != nil {
			return nil, errors.Wrap(err, op)
		}
		return &models.CustomSubscription{
			CustomSubscriptionID: cumSubID,
			Title:                title,
			Description:          description,
			Cost:                 cost,
			Layer:                layer,
		}, nil
	}
	return nil, nil
}
