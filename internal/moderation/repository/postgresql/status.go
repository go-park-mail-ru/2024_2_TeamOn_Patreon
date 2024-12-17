package postgresql

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/moderation/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
	"time"
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

	// getPostsByStatus
	// Input: $1 - status, $2 - limit, $3 - offset
	// Output: postID, title, description, authorID, authorName, createdAt
	getPostsByStatus = `
		select post_id, title, about,  user_id, username, created_date
		from Post
			join people using (user_id)
		where post_status_id = (select post_status_id from post_status where status = $1) 
		ORDER BY 
		    created_date DESC
		limit $2
		offset $3;
`

	// updateStatusOfPost - обновляет статус поста
	// Input: $1 - postID, $2 - status
	// Output: empty
	updateStatusOfPost = `
		update Post
		SET post_status_id = (select post_status_id from Post_Status where status = $2)
		WHERE post_id = $1
`

	// getStatusByPost
	// Input: $1 - PostID
	// Output: status
	getStatusByPost = `
	select status
	from Post
		join post_status using (post_status_id)
	where post_id = $1
`
)

func (mr *ModerationRepository) GetPostLayerByPostID(ctx context.Context, postID string) (int, error) {
	op := "moderation.repository.GetPostLayerByPostID"

	rows, err := mr.db.Query(ctx, getPostLayerByPostIdSQL, postID)
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

func (mr *ModerationRepository) UpdatePostStatus(ctx context.Context, postID string, status string) error {
	op := "moderation.repository.post.UpdatePostStatus"

	_, err := mr.db.Exec(ctx, updateStatusOfPost, postID, status)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (mr *ModerationRepository) GetPostsByStatus(ctx context.Context, status string, limit, offset int) ([]*models.Post, error) {
	op := "moderation.repository.GetPostsByStatus"

	rows, err := mr.db.Query(ctx, getPostsByStatus, status, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		postID         string
		title          string
		content        string
		authorId       string
		authorUsername string
		createdDate    time.Time
	)

	posts := make([]*models.Post, 0)

	for rows.Next() {
		if err = rows.Scan(&postID, &title, &content, &authorId, &authorUsername, &createdDate); err != nil {
			return nil, errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op,
			"Got  post: post_id=%v title=%v authorId=%v authorUsername=%v status=%v created_date=%v",
			postID, title, authorId, authorUsername, status, createdDate)
		posts = append(posts, &models.Post{
			PostID:         postID,
			Title:          title,
			Content:        content,
			AuthorID:       authorId,
			AuthorUsername: authorUsername,
			Status:         status,
			CreatedAt:      createdDate,
		})

	}

	return posts, nil
}

func (mr *ModerationRepository) GetStatusByPostID(ctx context.Context, postID string) (string, error) {
	op := "internal.content.repository.subscription.CheckCustomLayer"

	rows, err := mr.db.Query(ctx, getStatusByPost, postID)
	if err != nil {
		return "", errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		status string
	)

	for rows.Next() {
		if err = rows.Scan(&status); err != nil {
			return "", errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op, "Got  layer='%v' for status='%v'", status, postID)
		return status, nil
	}

	return "", nil
}
