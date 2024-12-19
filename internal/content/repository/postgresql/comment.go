package postgresql

import (
	"context"
	"time"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
)

const (
	// insertCommentSQL - сохраняет 1 коммент
	// Input: $1 commentID, $2 postID, $3 userID, $4 content
	// Output: empty
	insertCommentSQL = `
		INSERT INTO Comment (comment_id, post_id, user_id, about, created_date) VALUES
    ($1, $2, $3, $4, NOW())
  `

	// deleteCommentSQL - удаляет 1 коммент по id
	// Input: $1 commentID
	// Output: empty
	deleteCommentSQL = `
		delete from Comment where comment_id = $1
`
	// updateContentOfComment - обновляет коммент
	// Input: $1 - commentID, $2 - content
	// Output: empty
	updateContentOfComment = `
		update Comment
		SET about = $2, updated_date = NOW()
		WHERE comment_id = $1
`

	// getAuthorOfComment - возвращает автора комента
	// Input: $1 comment_id
	// Output: user_id
	getAuthorOfComment = `
		select user_id
		from People
		JOIN Comment USING (user_id)
		where comment_id = $1
`

	// getCommentsByPostID - возвращает комментарии к посту
	// Input: $1 - postID, $2 - limit, $3 - offset
	// Output: [commentID, userID, content, username, created_date]
	getCommentsByPostID = `
		select 
			comment.comment_id,
			comment.user_id,
			comment.about,
			people.username,
			comment.created_date
		from comment
			join people using (user_id)
		where comment.post_id = $1
		order by comment.created_date
		limit $2
		offset $3
`
)

func (cr *ContentRepository) CreateComment(ctx context.Context, userID, postID, commentID string, content string) error {
	op := "content.repository.CreateComment"

	_, err := cr.db.Exec(ctx, insertCommentSQL,
		commentID, postID, userID, content)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

func (cr *ContentRepository) UpdateComment(ctx context.Context, commentID string, content string) error {
	op := "content.repository.CreateComment"

	_, err := cr.db.Exec(ctx, updateContentOfComment,
		commentID, content)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

func (cr *ContentRepository) DeleteComment(ctx context.Context, commentID string) error {
	op := "content.repository.CreateComment"

	_, err := cr.db.Exec(ctx, deleteCommentSQL,
		commentID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

func (cr *ContentRepository) GetUserIDByCommentID(ctx context.Context, commentID string) (string, error) {
	op := "content.repository.post.GetUserIDByCommentID"

	rows, err := cr.db.Query(ctx, getAuthorOfComment, commentID)
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
		logger.StandardDebugF(ctx, op, "Got author='%s' of comment='%v'", authorID, commentID)
		return authorID, nil
	}
	return "", errors.Wrap(global.ErrCommentDoesntExist, op)
}

func (cr *ContentRepository) GetCommentsByPostID(ctx context.Context, postID string, limit, offset int) ([]*models.Comment, error) {
	op := "internal.content.repository.postgresql.GetPopularPostsForUser"

	rows, err := cr.db.Query(ctx, getCommentsByPostID, postID, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		commentID   string
		userID      string
		content     string
		username    string
		createdDate time.Time
	)

	comments := make([]*models.Comment, 0)

	for rows.Next() {
		if err = rows.Scan(&commentID, &userID, &content, &username, &createdDate); err != nil {
			return nil, errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op,
			"Got  comment: comment_id=%v content=%v userID=%v username=%v created_date=%v",
			commentID, content, userID, username, createdDate)
		comments = append(comments, &models.Comment{
			CommentID: commentID,
			Content:   content,
			UserID:    userID,
			Username:  username,
			CreatedAt: createdDate,
		})

	}

	return comments, nil
}
