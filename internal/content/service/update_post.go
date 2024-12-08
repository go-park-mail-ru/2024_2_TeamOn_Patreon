package service

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/service/validate"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (b *Behavior) UpdatePost(ctx context.Context, userID string, postID string, title string, about string) error {
	op := "content.service.behavior.UpdatePost"

	title, about, _, err := validate.Post(ctx, title, about, 0)
	if err != nil {
		return errors.Wrap(err, op)
	}

	if ok := utils.IsValidUUIDv4(userID); !ok {
		return errors.Wrap(global.ErrIsInvalidUUID, op)
	}

	isAuthor, err := b.isUserAuthorOfPost(ctx, postID, userID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	if !isAuthor {
		return errors.Wrap(global.ErrNotEnoughRights, op)
	}

	err = b.updateTitleInPost(ctx, postID, title)
	if err != nil {
		return errors.Wrap(err, op)
	}

	err = b.updateContentInPost(ctx, postID, about)
	if err != nil {
		return errors.Wrap(err, op)
	}

	go b.moderatePost(ctx, postID, title, about, true)

	return nil
}

func (b *Behavior) updateTitleInPost(ctx context.Context, postID string, title string) error {
	op := "content.service.behavior.updateTitleInPost"

	if title == "" {
		return nil
	}

	logger.StandardDebugF(ctx, op, "Want to update title=%v of post with id: %s", postID, title)
	err := b.rep.UpdateTitleOfPost(ctx, postID, title)
	if err != nil {
		return errors.Wrap(err, op)
	}
	logger.StandardDebugF(ctx, op, "Update title of post with id: %s", postID)
	return nil
}

func (b *Behavior) updateContentInPost(ctx context.Context, postID string, content string) error {
	op := "content.service.behavior.updateContentInPost"

	if content == "" {
		return nil
	}

	logger.StandardDebugF(ctx, op, "Want to update content=%v of post with id: %s", postID, content)
	err := b.rep.UpdateContentOfPost(ctx, postID, content)
	if err != nil {
		return errors.Wrap(err, op)
	}
	logger.StandardDebugF(ctx, op, "Update content in post with id: %v", postID)
	return nil
}
