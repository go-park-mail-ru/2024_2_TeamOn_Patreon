package postgresql

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

const (
	// getPostLayerBuPostIdSQL - возвращает уровень поста по его ид
	// Input: $1 postId
	// Output: layer (int) - минимальный, уровень подписки, на котором можно смотреть пост
	getPostLayerBuPostIdSQL = `
		select layer
			from Post
			join Subscription_Layer USING (subscription_layer_id)
		where post_id = $1;
`
)

func (cr *ContentRepository) GetPostLayerBuPostId(ctx context.Context, postID uuid.UUID) (int, error) {
	op := "internal.content.repository.subscription.CheckCustomLayer"

	rows, err := cr.db.Query(ctx, getPostLayerBuPostIdSQL, postID)
	if err != nil {
		return 0, errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		layer int
	)

	for rows.Next() {
		if err = rows.Scan(&layer); err != nil {
			return 0, errors.Wrap(err, op)
		}
		logger.StandardDebugF(op, "Got  layer='%v' for post='%v'", layer, postID)
		return layer, nil
	}

	return 0, nil
}
