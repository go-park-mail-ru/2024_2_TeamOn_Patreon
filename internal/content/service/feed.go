package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	models2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

// GetPopularPostsForUser отдаем популярные посты для конкретного пользователя
func (b *Behavior) GetPopularPosts(ctx context.Context, userId string, opt *models2.FeedOpt) ([]*models.Post, error) {
	op := "service.behavior.GetPopularPosts"

	userIdUuid, err := uuid.FromString(userId)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	if userId == "" {
		b.getPopularPostsForAnon(ctx, opt.Offset, opt.Limit)
	}

	posts, err := b.getPopularPostsForUser(ctx, userIdUuid, opt.Offset, opt.Limit)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return posts, nil
}

// getPopularPostsForAnon отдаем популярные посты для конкретного пользователя
func (b *Behavior) getPopularPostsForAnon(ctx context.Context, offset, limits int) ([]models.Post, error) {
	op := "service.behavior.GetPopularPostsForAnon"

	posts, err := b.rep.GetPopularPosts(ctx, offset, limits)
	if err != nil {
		return nil, errors.Wrap(global.ErrServer, op)
	}
	return posts, nil
}

func (b *Behavior) GetFeedSubscriptionGet(userId string, opt *models2.FeedOpt) ([]*models.Post, error) {
	return nil, nil
}

func (b *Behavior) getPopularPostsForUser(ctx context.Context, UserId uuid.UUID, offset, limits int) ([]*models.Post, error) {
	op := "service.behavior.GetPopularPosts"

	posts, err := b.rep.GetPopularPostsForUser(ctx, UserId, offset, limits)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	err = b.rep.GetIsLikedForPosts(ctx, UserId, posts)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return posts, nil
}
