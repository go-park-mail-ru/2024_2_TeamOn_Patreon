package postgresql

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
)

const (
	// insertPostSQL - сохраняет 1 пост
	// Input: $1 postId, $2 userId, $3 title, $4 about, $5 layer - уровень подписки
	// Output: empty
	insertPostSQL = `
		INSERT INTO Post (post_id, user_id, title, about, subscription_layer_id, post_status_id, created_date) VALUES
    ($1, $2, $3, $4, (SELECT subscription_layer_id FROM Subscription_Layer WHERE layer = $5), 
     (SELECT post_status_id FROM Post_Status WHERE status = 'PUBLISHED'),
     NOW())
  `

	// deletePostSQL - удаляет 1 пост по id
	// Input: $1 postId
	// Output: empty
	deletePostSQL = `
		delete from Post where post_id = $1
`

	// getAuthorOfPost - возвращает id автора поста
	// Input: $1 postId
	// Output: authorId
	getAuthorOfPost = `
		select user_id
		from Post
		where post_id = $1
`

	// Update

	// updateTitleOfPost - обновляет название поста
	// Input: $1 - postId, $2 - title
	// Output: empty
	updateTitleOfPost = `
		update Post
		SET title = $2, updated_date = NOW(), post_status_id = (select post_status_id from Post_Status where status = 'PUBLISHED')
		WHERE post_id = $1
`

	// updateContentOfPost
	// Input: $1 - postId, $2 - content
	// Output: empty
	updateContentOfPost = `
		update Post
		SET about = $2, updated_date = NOW(), post_status_id = (select post_status_id from Post_Status where status = 'PUBLISHED')
		WHERE post_id = $1
`
)

func (cr *ContentRepository) InsertPost(ctx context.Context, userID string, postID string, title string, content string, layer int) error {
	op := "internal.content.repository.post.InsertPost"

	_, err := cr.db.Exec(ctx, insertPostSQL,
		postID, userID, title, content, layer)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

func (cr *ContentRepository) DeletePost(ctx context.Context, postID string) error {
	op := "internal.content.repository.post.DeletePost"

	_, err := cr.db.Exec(ctx, deletePostSQL, postID)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (cr *ContentRepository) GetAuthorOfPost(ctx context.Context, postID string) (string, error) {
	op := "internal.content.repository.post.GetAuthorOfPost"

	rows, err := cr.db.Query(ctx, getAuthorOfPost, postID)
	if err != nil {
		return "", errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		authorID string
	)

	for rows.Next() {
		if err = rows.Scan(&authorID); err != nil {
			return "", errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op, "Got author='%s' of post='%v'", authorID, postID)
		return authorID, nil
	}
	return "", errors.Wrap(global.ErrPostDoesntExists, op)
}

func (cr *ContentRepository) UpdateTitleOfPost(ctx context.Context, postID string, title string) error {
	op := "internal.content.repository.post.UpdateTitleOfPost"

	_, err := cr.db.Exec(ctx, updateTitleOfPost, postID, title)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (cr *ContentRepository) UpdateContentOfPost(ctx context.Context, postID string, content string) error {
	op := "internal.content.repository.post.UpdateContentOfPost"

	_, err := cr.db.Exec(ctx, updateContentOfPost, postID, content)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}
