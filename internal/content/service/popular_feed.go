package service

import (
	"context"
	pkgModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/pkg/errors"
)

// GetPopularPosts отдаем популярные посты для конкретного пользователя
func (b *Behavior) GetPopularPosts(ctx context.Context, userID string, opt *models.FeedOpt) ([]*pkgModels.Post, error) {
	op := "service.behavior.GetPopularPosts"

	if userID == "" {
		posts, err := b.getPopularPostsForAnon(ctx, opt.Offset, opt.Limit)
		if err != nil {
			return nil, errors.Wrap(err, op)
		}
		return posts, nil
	}

	if ok := utils.IsValidUUIDv4(userID); !ok {
		return nil, errors.Wrap(global.ErrIsInvalidUUID, op)
	}

	posts, err := b.getPopularPostsForUser(ctx, userID, opt.Offset, opt.Limit)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return posts, nil
}

// getPopularPostsForAnon отдаем популярные посты для конкретного пользователя
func (b *Behavior) getPopularPostsForAnon(ctx context.Context, offset, limits int) ([]*pkgModels.Post, error) {
	op := "service.behavior.GetPopularPostsForAnon"

	posts, err := b.rep.GetPopularPosts(ctx, offset, limits)
	if err != nil {
		return nil, errors.Wrap(global.ErrServer, op)
	}
	return posts, nil
}

func (b *Behavior) getPopularPostsForUser(ctx context.Context, userID string, offset, limits int) ([]*pkgModels.Post, error) {
	op := "service.behavior.GetPopularPosts"

	posts, err := b.rep.GetPopularPostsForUser(ctx, userID, offset, limits)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	err = b.rep.GetIsLikedForPosts(ctx, userID, posts)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return posts, nil
}
