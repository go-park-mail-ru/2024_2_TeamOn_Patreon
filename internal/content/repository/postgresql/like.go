package postgresql

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
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

	// sendNotificationLike - сохраняет новое уведомление о лайке поста
	// Input: $1 notification_id, $2 - user_id, $3 - sender_id, $4 - about
	// Output: empty
	sendNotificationLike = `
		INSERT INTO 
			notification (notification_id, user_id, sender_id, about)
		VALUES
			($1, $2, $3, $4);
`
)

func (cr *ContentRepository) GetPostLikeID(ctx context.Context, userID string, postID string) (string, error) {
	op := "internal.content.repository.postgresql.GetPostLikeId"

	rows, err := cr.db.Query(ctx, getPostLikeIdSQL, postID, userID)
	if err != nil {
		return "", errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		postLikeID string
	)

	for rows.Next() {
		if err = rows.Scan(&postLikeID); err != nil {
			return "", errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op, "Got  postLikeID='%v' for post='%v'", postLikeID, postID)
		return postLikeID, nil
	}

	return "", nil
}

func (cr *ContentRepository) InsertLikePost(ctx context.Context, userID string, postID string) error {
	op := "internal.content.repository.InsertLikePost"

	_, err := cr.db.Exec(ctx, insertPostLikeSQL, postID, userID)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (cr *ContentRepository) DeleteLikePost(ctx context.Context, userID string, postID string) error {
	op := "internal.content.repository.DeleteLikePost"

	_, err := cr.db.Exec(ctx, deletePostLikeSQL, postID, userID)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (cr *ContentRepository) GetPostLikes(ctx context.Context, postID string) (int, error) {
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
		logger.StandardDebugF(ctx, op, "Got  likes='%v' for post='%v'", likes, postID)
		return likes, nil
	}

	return 0, nil
}

// GetUsername - получение имени пользователя по userID
func (cr *ContentRepository) GetUsername(ctx context.Context, userID string) (string, error) {
	op := "internal.account.repository.GetUsername"

	query := `
		SELECT 
			username
		FROM
			people
		WHERE
			user_id = $1;
	`

	rows, err := cr.db.Query(ctx, query, userID)
	if err != nil {
		return "", errors.Wrap(err, op)
	}

	defer rows.Close()

	var username string

	for rows.Next() {
		if err = rows.Scan(&username); err != nil {
			return "", errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op, "Got username='%v' for userID='%v'", username, userID)
		return username, nil
	}
	return "", nil
}

func (cr *ContentRepository) SendNotificationOfLike(ctx context.Context, message, userID, authorID string) error {
	op := "internal.content.repository.postgresql.SendNotificationOfLike"

	notificationID := utils.GenerateUUID()

	// Не путать: userID ($2) = authorID (получатель), senderID ($3) = userID (отправитель)
	_, err := cr.db.Exec(ctx, sendNotificationLike, notificationID, authorID, userID, message)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}
