package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	models2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

func (b *Behavior) GetFeedSubscription(ctx context.Context, userId string, opt *models2.FeedOpt) ([]*models.Post, error) {
	op := "internal.content.service.subscription_feed.GetFeedSubscription"

	if userId == "" {
		return nil, errors.Wrap(global.ErrUserNotAuthorized, op)
	}

	userIdUuid, err := uuid.FromString(userId)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	posts, err := b.getSubscriptionPostsForUser(ctx, userIdUuid, opt.Offset, opt.Limit)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return posts, nil
}

func (b *Behavior) getSubscriptionPostsForUser(ctx context.Context, UserId uuid.UUID, offset, limits int) ([]*models.Post, error) {
	op := "service.behavior.getSubscriptionPostsForUser"

	posts, err := b.rep.GetSubscriptionPostsForUser(ctx, UserId, offset, limits)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	err = b.rep.GetIsLikedForPosts(ctx, UserId, posts)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return posts, nil
}
