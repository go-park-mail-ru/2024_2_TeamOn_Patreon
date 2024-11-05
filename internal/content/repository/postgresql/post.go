package postgresql

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

const (
	// insertPostSQL - сохраняет 1 пост
	// Input: $1 postId, $2 userId, $3 title, $4 about, $5 layer - уровень подписки
	// Output: empty
	insertPostSQL = `
		INSERT INTO Post (post_id, user_id, title, about, subscription_layer_id, created_date) VALUES
    ($1, $2, $3, $4, (SELECT subscription_layer_id FROM Subscription_Layer WHERE layer = $5), NOW())
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
		SET title = $2
		WHERE post_id = $1
`

	// updateContentOfPost
	// Input: $1 - postId, $2 - content
	// Output: empty
	updateContentOfPost = `
		update Post
		SET about = $2
		WHERE post_id = $1
`
)

func (cr *ContentRepository) InsertPost(ctx context.Context, userId uuid.UUID, postId uuid.UUID, title string, content string, layer int) error {
	op := "internal.content.repository.post.InsertPost"

	_, err := cr.db.Exec(ctx, insertPostSQL,
		postId, userId, title, content, layer)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

func (cr *ContentRepository) DeletePost(ctx context.Context, postID uuid.UUID) error {
	op := "internal.content.repository.post.DeletePost"

	_, err := cr.db.Exec(ctx, deletePostSQL, postID)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (cr *ContentRepository) GetAuthorOfPost(ctx context.Context, postID uuid.UUID) (uuid.UUID, error) {
	op := "internal.content.repository.post.GetAuthorOfPost"

	rows, err := cr.db.Query(ctx, getAuthorOfPost, postID)
	if err != nil {
		return uuid.UUID{}, errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		authorId uuid.UUID
	)

	for rows.Next() {
		if err = rows.Scan(&authorId); err != nil {
			return uuid.UUID{}, errors.Wrap(err, op)
		}
		logger.StandardDebugF(op, "Got author='%s' of post='%v'", authorId, postID)
		return authorId, nil
	}
	return uuid.UUID{}, errors.Wrap(global.ErrPostDoesntExists, op)
}

func (cr *ContentRepository) UpdateTitleOfPost(ctx context.Context, postID uuid.UUID, title string) error {
	op := "internal.content.repository.post.UpdateTitleOfPost"

	_, err := cr.db.Exec(ctx, updateTitleOfPost, postID, title)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (cr *ContentRepository) UpdateContentOfPost(ctx context.Context, postID uuid.UUID, content string) error {
	op := "internal.content.repository.post.UpdateContentOfPost"

	_, err := cr.db.Exec(ctx, updateContentOfPost, postID, content)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}
