package service

import (
	"context"
	pkgModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/pkg/errors"
)

func (b *Behavior) GetFeedSubscription(ctx context.Context, userID string, opt *models.FeedOpt) ([]*pkgModels.Post, error) {
	op := "internal.content.service.subscription_feed.GetFeedSubscription"

	if userID == "" {
		return nil, errors.Wrap(global.ErrUserNotAuthorized, op)
	}

	if ok := utils.IsValidUUIDv4(userID); !ok {
		return nil, errors.Wrap(global.ErrIsInvalidUUID, op)
	}

	posts, err := b.getSubscriptionPostsForUser(ctx, userID, opt.Offset, opt.Limit)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return posts, nil
}

func (b *Behavior) getSubscriptionPostsForUser(ctx context.Context, userID string, offset, limits int) ([]*pkgModels.Post, error) {
	op := "service.behavior.getSubscriptionPostsForUser"

	posts, err := b.rep.GetSubscriptionPostsForUser(ctx, userID, offset, limits)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	err = b.rep.GetIsLikedForPosts(ctx, userID, posts)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return posts, nil
}
