package service

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	models2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/pkg/errors"
)

// GetPopularPostsForUser отдаем популярные посты для конкретного пользователя
func (b *Behavior) GetPopularPostsForUser(userId string, opt *models2.FeedOpt) ([]models.Post, error) {
	op := "service.behavior.GetPopularPostsForUser"

	posts, err := b.rep.GetPopularPosts(opt.Offset, opt.Limit)
	if err != nil {
		return nil, errors.Wrap(global.ErrServer, op)
	}
	return posts, nil
}

// GetPopularPostsForAnon отдаем популярные посты для конкретного пользователя
func (b *Behavior) GetPopularPostsForAnon(opt *models2.FeedOpt) ([]models.Post, error) {
	op := "service.behavior.GetPopularPostsForAnon"

	posts, err := b.rep.GetPopularPosts(opt.Offset, opt.Limit)
	if err != nil {
		return nil, errors.Wrap(global.ErrServer, op)
	}
	return posts, nil
}

func (b *Behavior) GetFeedSubscriptionGet(userId string, opt *models2.FeedOpt) ([]models.Post, error) {
	return nil, nil
}
