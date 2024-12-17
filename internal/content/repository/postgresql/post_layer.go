package postgresql

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
)

const (
	// getPostLayerByPostIdSQL - возвращает уровень поста по его ид
	// Input: $1 postId
	// Output: layer (int) - минимальный, уровень подписки, на котором можно смотреть пост
	getPostLayerByPostIdSQL = `
		select layer
			from Post
			join Subscription_Layer USING (subscription_layer_id)
		where post_id = $1;
`
)

func (cr *ContentRepository) GetPostLayerByPostID(ctx context.Context, postID string) (int, error) {
	op := "internal.content.repository.subscription.GetPostLayerByPostID"

	rows, err := cr.db.Query(ctx, getPostLayerByPostIdSQL, postID)
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
		logger.StandardDebugF(ctx, op, "Got  layer='%v' for post='%v'", layer, postID)
		return layer, nil
	}

	return 0, nil
}
