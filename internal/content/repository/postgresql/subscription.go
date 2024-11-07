package postgresql

import (
	"context"
	"time"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
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

	// getSubscriptionFeedForUser возвращает посты отсортированные по дате по убывающей, которые
	// может смотреть пользователь среди всех постов авторов, на которых он подписан
	// Output: postID, title, about, authorId, authorUsername, likes, created_date
	// likes - количество лайков
	// Input: получает $1 userId - uuid пользователя, {$2 offset} и { $3 limit}
	getSubscriptionFeedForUser = `
SELECT 
    post.post_id, 
    post.Title, 
    post.About, 
    author.user_id AS author_id, 
    author.Username AS author_username, 
    (SELECT like_post_id FROM Like_Post where like_post.post_id = post.post_id) AS likes,
    post.created_date
FROM 
    post
JOIN 
    People AS author ON author.user_id = post.user_id 
RIGHT OUTER JOIN 
    Subscription_Layer ON Subscription_Layer.subscription_layer_id = post.subscription_layer_id
LEFT OUTER JOIN 
    Like_Post USING (post_id)
WHERE 
    Subscription_Layer.layer <= (
        SELECT COALESCE(Subscription_Layer.layer, 0)
        FROM Subscription
        JOIN Custom_Subscription ON Subscription.custom_subscription_id = Custom_Subscription.custom_subscription_id
        JOIN Subscription_Layer ON Custom_Subscription.subscription_layer_id = Subscription_Layer.subscription_layer_id
        WHERE Custom_Subscription.author_id = author.user_id AND Subscription.user_id = $1
    )
ORDER BY 
    created_date DESC
LIMIT $3
OFFSET $2
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
		logger.StandardDebugF(ctx, op, "Got subscription='%s' user='%s' layer='%v' is='%v'",
			subscription, authorId, layer, layerExists)
		return layerExists, nil
	}

	return false, nil
}

func (cr ContentRepository) GetSubscriptionPostsForUser(ctx context.Context, userId uuid.UUID, offset int, limits int) ([]*models.Post, error) {
	op := "internal.content.repository.subscription.GetSubscriptionPostsForUser"

	rows, err := cr.db.Query(ctx, getSubscriptionFeedForUser, userId, offset, limits)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		postID         uuid.UUID
		title          string
		content        string
		authorId       uuid.UUID
		authorUsername string
		likes          int
		createdDate    time.Time
	)

	posts := make([]*models.Post, 0)

	for rows.Next() {
		if err = rows.Scan(&postID, &title, &content, &authorId, &authorUsername, &likes, &createdDate); err != nil {
			return nil, errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op,
			"Got  post: post_id=%v title=%v authorId=%v authorUsername=%v likes=%v created_date=%v",
			postID, title, authorId, authorUsername, likes, createdDate)
		posts = append(posts, &models.Post{
			PostId:         postID.String(),
			Title:          title,
			Content:        content,
			AuthorId:       authorId.String(),
			AuthorUsername: authorUsername,
			Likes:          likes,
			CreatedDate:    createdDate,
		})

	}

	return posts, nil
}
