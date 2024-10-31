package postgresql

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

const (
	// getCustomSubscriptionByUserIdAndLayerSQl - возвращает имя кастомной подписки у автора на определенном уровне
	// Input: $1 - userId, $2 - уровень подписки (int)
	// Output: custom_name - кастомное имя подписки
	getCustomSubscriptionByUserIdAndLayerSQl = `
		select custom_name
		from
		    Custom_Subscription
			join Subscription_Layer USING (subscription_layer_id)
		where 
		    Custom_Subscription.author_id = $1
			and Subscription_Layer.layer = $2
`
)

// subscription

func (cr *ContentRepository) CheckCustomLayer(ctx context.Context, authorId uuid.UUID, layer int) (bool, error) {
	op := "internal.content.repository.subscription.CheckCustomLayer"

	rows, err := cr.db.Query(ctx, getCustomSubscriptionByUserIdAndLayerSQl, authorId, layer)
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		subscription string
	)

	for rows.Next() {
		if err = rows.Scan(&subscription); err != nil {
			return false, errors.Wrap(err, op)
		}
		layerExists := subscription != ""
		logger.StandardDebugF(op, "Got subscription='%s' user='%s' layer='%v' is='%v'",
			subscription, authorId, layer, layerExists)
		return layerExists, nil
	}

	return false, nil
}
