package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"

	pkgModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/pkg/errors"
)

// GetAuthorPosts лента постов на странице автора
// userId - пустой - не авторизован - лайков нет
// userId - невалидный - ошибка
// authorId - me - страница автора
// authorId - невалидный ошибка
func (b *Behavior) GetAuthorPosts(ctx context.Context, userID, authorID string, opt *models.FeedOpt) ([]*pkgModels.Post, error) {
	op := "internal.content.repository.author_feed.GetAuthorPosts"

	// Нельзя посмотреть свою страницу анониму
	if userID == "" && authorID == "me" {
		return nil, errors.Wrap(global.ErrUserNotAuthorized, op)
	}

	// аноним смотрит посты автора
	if userID == "" {
		if ok := utils.IsValidUUIDv4(authorID); ok != true {
			return nil, errors.Wrap(global.ErrIsInvalidUUID, op)
		}

		posts, err := b.getAuthorPostsForAnon(ctx, authorID, opt.Offset, opt.Limit)
		if err != nil {
			return nil, errors.Wrap(err, op)
		}
		return posts, nil
	}

	// юзер
	if ok := utils.IsValidUUIDv4(userID); !ok {
		return nil, errors.Wrap(global.ErrIsInvalidUUID, op)
	}

	// автор смотрит свои посты
	if authorID == "me" || userID == authorID {
		isAuthor, err := b.isUserAuthor(ctx, userID)
		if err != nil {
			return nil, errors.Wrap(err, op)
		}

		// не у автора нет постов
		if !isAuthor {
			// Optional: заменить на специализированную ошибку
			return nil, errors.Wrap(global.ErrNotEnoughRights, op)
		}

		posts, err := b.getAuthorMyPosts(ctx, userID, opt.Offset, opt.Limit)
		if err != nil {
			return nil, errors.Wrap(err, op)
		}
		return posts, nil
	}

	if ok := utils.IsValidUUIDv4(authorID); !ok {
		return nil, errors.Wrap(global.ErrIsInvalidUUID, op)
	}

	// юзер смотрит посты автора
	//  может быть как подписчиком, так и неи
	posts, err := b.getAuthorPostsForUser(ctx, userID, authorID, opt.Offset, opt.Limit)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	return posts, nil
}

func (b *Behavior) getAuthorMyPosts(ctx context.Context, userID string, offset, limit int) ([]*pkgModels.Post, error) {
	op := "internal.content.repository.author_post.GetAuthorMyPosts"

	logger.StandardDebugF(ctx, op, "Getting my posts for authorId='%v'", userID)

	posts, err := b.rep.GetAuthorPostsForMe(ctx, userID, offset, limit)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	err = b.getLikePostsForUser(ctx, userID, posts)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return posts, nil
}

func (b *Behavior) getAuthorPostsForAnon(ctx context.Context, authorID string, offset, limit int) ([]*pkgModels.Post, error) {
	op := "internal.content.repository.author_feed.GetAuthorPostsForAnon"

	logger.StandardDebugF(ctx, op, "Getting posts by authorID='%v' for anon", authorID)

	posts, err := b.rep.GetAuthorPostsForAnon(ctx, authorID, offset, limit)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	return posts, nil
}

func (b *Behavior) getAuthorPostsForUser(ctx context.Context, userID, authorID string, offset, limit int) ([]*pkgModels.Post, error) {
	op := "internal.content.repository.author_post.GetAuthorPostsForUser"

	logger.StandardDebugF(ctx, op, "Getting posts by authorID='%v' for userID='%v'", authorID, userID)

	layer, err := b.rep.GetUserLayerOfAuthor(ctx, userID, authorID)

	posts, err := b.rep.GetAuthorPostsForLayer(ctx, layer, authorID, offset, limit)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	err = b.getLikePostsForUser(ctx, userID, posts)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return posts, nil
}
