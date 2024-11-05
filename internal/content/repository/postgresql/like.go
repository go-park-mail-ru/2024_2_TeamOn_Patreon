package postgresql

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

const (
	// insertPostLikeSQL - добавляет лайк на пост
	// Input: $1 post_id, $2 - user_id
	// Output: empty
	insertPostLikeSQL = `
INSERT INTO Like_Post (like_post_id, post_id, user_id, posted_date) VALUES
    (gen_random_uuid(), $1, $2, NOW())
	;
`

	// deletePostLikeSQL - удаляет лайк пользователя на пост
	// Input: $1 post_id, $2 - user_id
	// Output: empty
	deletePostLikeSQL = `delete from Like_Post where user_id = $2 and post_id = $1;`

	// getCountPostLikesSQL - возвращает количество лайков поста
	// Input: $1 postId (uuid)
	// Output: likes (int)
	getCountPostLikesSQL = `
	SELECT COUNT(*) as likes
	FROM Like_Post
	JOIN Post USING (post_id)
	WHERE post_id = $1;
`

	// getPostLikeIdSQL - возвращает id сущности лайка на пост
	// Input: $1 post_id, $2 - user_id
	// Output: empty
	getPostLikeIdSQL = `
		select like_post_id
		from Like_Post
		where post_id = $1 and user_id = $2;
`
)

func (cr *ContentRepository) GetPostLikeId(ctx context.Context, userId uuid.UUID, postID uuid.UUID) (uuid.UUID, error) {
	op := "internal.content.repository.postgresql.GetPostLikeId"

	rows, err := cr.db.Query(ctx, getPostLikeIdSQL, postID, userId)
	if err != nil {
		return uuid.UUID{}, errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		postLikeId uuid.UUID
	)

	for rows.Next() {
		if err = rows.Scan(&postLikeId); err != nil {
			return uuid.UUID{}, errors.Wrap(err, op)
		}
		logger.StandardDebugF(op, "Got  postLikeId='%v' for post='%v'", postLikeId, postID)
		return postLikeId, nil
	}

	return uuid.UUID{}, nil
}

func (cr *ContentRepository) InsertLikePost(ctx context.Context, userId uuid.UUID, postID uuid.UUID) error {
	op := "internal.content.repository.InsertLikePost"

	_, err := cr.db.Exec(ctx, insertPostLikeSQL, postID, userId)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (cr *ContentRepository) DeleteLikePost(ctx context.Context, userId uuid.UUID, postID uuid.UUID) error {
	op := "internal.content.repository.DeleteLikePost"

	_, err := cr.db.Exec(ctx, deletePostLikeSQL, postID, userId)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (cr *ContentRepository) GetPostLikes(ctx context.Context, postID uuid.UUID) (int, error) {
	op := "internal.content.repository.postgresql.GetPostLikes"

	rows, err := cr.db.Query(ctx, getCountPostLikesSQL, postID)
	if err != nil {
		return 0, errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		likes int
	)

	for rows.Next() {
		if err = rows.Scan(&likes); err != nil {
			return 0, errors.Wrap(err, op)
		}
		logger.StandardDebugF(op, "Got  likes='%v' for post='%v'", likes, postID)
		return likes, nil
	}

	return 0, nil
}
