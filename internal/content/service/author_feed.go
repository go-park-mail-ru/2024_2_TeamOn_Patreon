package service

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	models2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

// GetAuthorPosts лента постов на странице автора
// userId - пустой - не авторизован - лайков нет
// userId - невалидный - ошибка
// authorId - me - страница автора
// authorId - невалидный ошибка
func (b *Behavior) GetAuthorPosts(ctx context.Context, userId, authorId string, opt *models2.FeedOpt) ([]*models.Post, error) {
	op := "internal.content.repository.author_feed.GetAuthorPosts"

	// Нельзя посмотреть свою страницу анониму
	if userId == "" && authorId == "me" {
		return nil, errors.Wrap(global.ErrUserNotAuthorized, op)
	}

	// аноним смотрит посты автора
	if userId == "" {
		authorIdUuid, err := uuid.FromString(authorId)
		if err != nil {
			return nil, errors.Wrap(err, op)
		}

		posts, err := b.getAuthorPostsForAnon(ctx, authorIdUuid, opt.Offset, opt.Limit)
		if err != nil {
			return nil, errors.Wrap(err, op)
		}
		return posts, nil
	}

	// юзер
	userIdUuid, err := uuid.FromString(userId)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	// автор смотрит свои посты
	if authorId == "me" || userId == authorId {
		isAuthor, err := b.isUserAuthor(ctx, userIdUuid)
		if err != nil {
			return nil, errors.Wrap(err, op)
		}

		// не у автора нет постов
		if !isAuthor {
			// Optional: заменить на специализированную ошибку
			return nil, errors.Wrap(global.ErrNotEnoughRights, op)
		}

		posts, err := b.getAuthorMyPosts(ctx, userIdUuid, opt.Offset, opt.Limit)
		if err != nil {
			return nil, errors.Wrap(err, op)
		}
		return posts, nil
	}

	// достаем автора
	authorIdUuid, err := uuid.FromString(authorId)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	// юзер смотрит посты автора
	//  может быть как подписчиком, так и неи
	posts, err := b.getAuthorPostsForUser(ctx, userIdUuid, authorIdUuid, opt.Offset, opt.Limit)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	return posts, nil
}

func (b *Behavior) getAuthorMyPosts(ctx context.Context, userId uuid.UUID, offset, limit int) ([]*models.Post, error) {
	op := "internal.content.repository.author_post.GetAuthorMyPosts"

	logger.StandardDebugF(ctx, op, "Getting my posts for authorId='%v'", userId)

	posts, err := b.rep.GetAuthorPostsForMe(ctx, userId, offset, limit)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	err = b.getLikePostsForUser(ctx, userId, posts)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return posts, nil
}

func (b *Behavior) getAuthorPostsForAnon(ctx context.Context, authorId uuid.UUID, offset, limit int) ([]*models.Post, error) {
	op := "internal.content.repository.author_feed.GetAuthorPostsForAnon"

	logger.StandardDebugF(ctx, op, "Getting posts by authorId='%v' for anon", authorId)

	posts, err := b.rep.GetAuthorPostsForAnon(ctx, authorId, offset, limit)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	return posts, nil
}

func (b *Behavior) getAuthorPostsForUser(ctx context.Context, userId, authorId uuid.UUID, offset, limit int) ([]*models.Post, error) {
	op := "internal.content.repository.author_post.GetAuthorPostsForUser"

	logger.StandardDebugF(ctx, op, "Getting posts by authorId='%v' for userId='%v'", authorId, userId)

	layer, err := b.rep.GetUserLayerOfAuthor(ctx, userId, authorId)

	posts, err := b.rep.GetAuthorPostsForLayer(ctx, layer, authorId, offset, limit)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	err = b.getLikePostsForUser(ctx, userId, posts)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return posts, nil
}
