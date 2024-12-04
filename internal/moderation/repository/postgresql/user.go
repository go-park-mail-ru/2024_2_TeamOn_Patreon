package postgresql

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
)

const (
	// getAuthorOfPost - возвращает id автора поста
	// Input: $1 postId
	// Output: authorId
	getAuthorOfPost = `
		select user_id
		from Post
		where post_id = $1
`

	// getUserLayerOfAuthor - получение уровня подписки пользователя на определенного автора
	// Input: $1 - userId, $2 - authorId
	// Output: layer (int)
	getUserLayerOfAuthor = `
		select layer
		from Subscription
		join Custom_Subscription USING (custom_subscription_id)
		join Subscription_Layer ON Subscription_Layer.subscription_layer_id = Custom_Subscription.subscription_layer_id
		where Custom_Subscription.author_id = $2
		and Subscription.user_id = $1
		;
`

	// getUserRoleSQL - возвращает текстовое имя роли пользователя
	// Input: userId
	// Output: role_name ('Reader' or 'Author' or '')
	getUserRoleSQL = `
select 
	role_default_name
from 
	Role
	join People USING (role_id)
where People.user_id = $1
	`
)

func (mr *ModerationRepository) GetAuthorOfPost(ctx context.Context, postID string) (string, error) {
	op := "internal.content.repository.post.GetAuthorOfPost"

	rows, err := mr.db.Query(ctx, getAuthorOfPost, postID)
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

func (mr *ModerationRepository) GetUserLayerOfAuthor(ctx context.Context, userID, authorID string) (int, error) {
	op := "internal.content.repository.post.GetUserLayerOfAuthor"

	logger.StandardDebugF(ctx, op, "Want to get user layer userID=%v, author = %v", userID, authorID)

	rows, err := mr.db.Query(ctx, getUserLayerOfAuthor, userID, authorID)
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
		logger.StandardDebugF(ctx, op, "Got layer= %s user= %s", layer, userID)
		return layer, nil
	}

	return 0, nil
}

func (mr *ModerationRepository) GetUserRole(ctx context.Context, userID string) (string, error) {
	op := "moderation.repository.user.GetUserRole"

	logger.StandardDebugF(ctx, op, "Want to get user role userID=%v, db = %v", userID, mr.db)

	rows, err := mr.db.Query(ctx, getUserRoleSQL, userID)
	if err != nil {
		return "", errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		role string
	)
	for rows.Next() {
		if err = rows.Scan(&role); err != nil {
			return "", errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op, "Got layer= %s user= %s", role, userID)
		return role, nil
	}

	return "", nil
}
