package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

func (b *Behavior) UpdatePost(ctx context.Context, userID string, postId string, title string, about string) error {
	op := "content.service.behavior.UpdatePost"

	authorId, err := uuid.FromString(userID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	postIdUuid, err := uuid.FromString(postId)
	if err != nil {
		return errors.Wrap(err, op)
	}

	isAuthor, err := b.isUserAuthorOfPost(ctx, postIdUuid, authorId)
	if err != nil {
		return errors.Wrap(err, op)
	}

	if !isAuthor {
		return errors.Wrap(global.ErrNotEnoughRights, op)
	}

	err = b.updateTitleInPost(ctx, postIdUuid, title)
	if err != nil {
		return errors.Wrap(err, op)
	}

	err = b.updateContentInPost(ctx, postIdUuid, about)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

func (b *Behavior) updateTitleInPost(ctx context.Context, postId uuid.UUID, title string) error {
	op := "content.service.behavior.updateTitleInPost"

	if title == "" {
		return nil
	}

	logger.StandardDebugF(op, "Want to update title=%v of post with id: %s", postId, title)
	err := b.rep.UpdateTitleOfPost(ctx, postId, title)
	if err != nil {
		return errors.Wrap(err, op)
	}
	logger.StandardDebugF(op, "Update title of post with id: %s", postId)
	return nil
}

func (b *Behavior) updateContentInPost(ctx context.Context, postId uuid.UUID, content string) error {
	op := "content.service.behavior.updateContentInPost"

	if content == "" {
		return nil
	}

	logger.StandardDebugF(op, "Want to update content=%v of post with id: %s", postId, content)
	err := b.rep.UpdateContentOfPost(ctx, postId, content)
	if err != nil {
		return errors.Wrap(err, op)
	}
	logger.StandardDebugF(op, "Update content in post with id: %v", postId)
	return nil
}
