package service

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

func (b *Behavior) DeletePost(ctx context.Context, userId, postId string) error {
	op := "service.behavior.DeletePost"

	userIdUuid, err := uuid.FromString(userId)
	if err != nil {
		return errors.Wrap(err, op)
	}
	postIdUuid, err := uuid.FromString(postId)
	logger.StandardDebugF(ctx, op, "userId= %v post_id = '%v'", userId, postId)
	if err != nil {
		return errors.Wrap(err, op)
	}

	// Проверка что юзер - автор этого поста
	userIsAuthor, err := b.isUserAuthorOfPost(ctx, postIdUuid, userIdUuid)
	if err != nil {
		return errors.Wrap(err, op)
	}

	if !userIsAuthor {
		return errors.Wrap(global.ErrNotEnoughRights, op)
	}

	// Удаление поста
	err = b.deletePost(ctx, postIdUuid)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

// isUserAuthorPost проверяет что пользователь автор поста
func (b *Behavior) isUserAuthorOfPost(ctx context.Context, postId, userId uuid.UUID) (bool, error) {
	op := "service.behavior.isUserAuthorPost"

	authorId, err := b.rep.GetAuthorOfPost(ctx, postId)
	logger.StandardDebugF(ctx, op, "Got author=%v of post=%v", authorId, postId)

	if err != nil {
		return false, errors.Wrap(err, op)
	}

	if authorId != userId {
		return false, nil
	}

	return true, nil
}

func (b *Behavior) deletePost(ctx context.Context, postId uuid.UUID) error {
	op := "service.behavior.deletePost"

	err := b.rep.DeletePost(ctx, postId)
	if err != nil {
		return errors.Wrap(err, op)
	}
	logger.StandardDebugF(ctx, op, "Successfully deleted post=%v", postId)
	return nil
}
