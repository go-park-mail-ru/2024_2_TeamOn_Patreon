package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/pkg/errors"
)

func (b *Behavior) GetFile(ctx context.Context, userID string, postID string) ([]*models.Media, error) {
	op := "content.service.GetFile"

	canSee, err := b.userCanSeePost(ctx, userID, postID)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	if !canSee {
		err = global.ErrNotEnoughRights
		return nil, errors.Wrap(err, op)
	}

	medias, err := b.getContentsByPost(ctx, postID)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return medias, nil
}

func (b *Behavior) getContentsByPost(ctx context.Context, postID string) ([]*models.Media, error) {
	op := "content.service.getContentsByPost"

	medias, err := b.rep.GetContentsByPost(ctx, postID)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	return medias, nil
}
