package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (b *Behavior) DeletePost(ctx context.Context, userID, postID string) error {
	op := "service.behavior.DeletePost"

	if ok := utils.IsValidUUIDv4(userID); !ok {
		return errors.Wrap(global.ErrIsInvalidUUID, op)
	}

	if ok := utils.IsValidUUIDv4(postID); !ok {
		return errors.Wrap(global.ErrIsInvalidUUID, op)
	}

	logger.StandardDebugF(ctx, op, "userID= %v post_id = '%v'", userID, postID)

	// Проверка что юзер - автор этого поста
	userIsAuthor, err := b.isUserAuthorOfPost(ctx, postID, userID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	if !userIsAuthor {
		return errors.Wrap(global.ErrNotEnoughRights, op)
	}

	// Удаление поста
	err = b.deletePost(ctx, postID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

// isUserAuthorPost проверяет что пользователь автор поста
func (b *Behavior) isUserAuthorOfPost(ctx context.Context, postID, userID string) (bool, error) {
	op := "service.behavior.isUserAuthorPost"

	authorId, err := b.rep.GetAuthorOfPost(ctx, postID)
	logger.StandardDebugF(ctx, op, "Got author=%v of post=%v", authorId, postID)

	if err != nil {
		return false, errors.Wrap(err, op)
	}

	if authorId != userID {
		return false, nil
	}

	return true, nil
}

func (b *Behavior) deletePost(ctx context.Context, postID string) error {
	op := "service.behavior.deletePost"

	err := b.rep.DeletePost(ctx, postID)
	if err != nil {
		return errors.Wrap(err, op)
	}
	logger.StandardDebugF(ctx, op, "Successfully deleted post=%v", postID)
	return nil
}
