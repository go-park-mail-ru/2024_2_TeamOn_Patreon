package service

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

func (b *Behavior) CreatePost(ctx context.Context, authorId string, title string, content string, layer int) (string, error) {
	op := "service.behavior.CreatePost"

	logger.StandardDebugF(ctx, op, "Got authorId:=%v title:=%v content:=%v layer:=%v", authorId, title, content, layer)

	newPostId, err := uuid.NewV4()
	if err != nil {
		return "", errors.Wrap(global.ErrServer, op)
	}

	uuidUserId, err := uuid.FromString(authorId)
	if err != nil {
		return "", errors.Wrap(err, op)
	}

	// Проверяем является ли человек автором
	isAuthor, err := b.isUserAuthor(ctx, uuidUserId)
	if err != nil {
		return "", errors.Wrap(err, op)
	}
	if !isAuthor {
		// Почему-то ошибка выделяется красным, хотя все ок
		return "", errors.Wrap(global.ErrNotEnoughRights, op)
	}

	// Проверяем есть ли кастомные подписки у этого пользователя на этом уровне
	layerExist, err := b.checkLayerExist(ctx, uuidUserId, layer)
	if err != nil {
		return "", errors.Wrap(err, op)
	}
	if !layerExist {
		return "", errors.Wrap(global.ErrBadRequest, op)
	}

	// Вставляем пост
	err = b.insertPost(ctx, uuidUserId, newPostId, title, content, layer)
	if err != nil {
		return "", errors.Wrap(err, op)
	}
	return newPostId.String(), nil
}

func (b *Behavior) checkCustomLayer(ctx context.Context, authorId uuid.UUID, layer int) (bool, error) {
	op := "content.service.checkCustomLayer"

	layerExists, err := b.rep.CheckCustomLayer(ctx, authorId, layer)
	if err != nil {
		logger.StandardDebugF(ctx, op, "checkCustomLayer err=%v", err)
		return false, errors.Wrap(err, op)
	}
	return layerExists, nil
}

func (b *Behavior) insertPost(ctx context.Context, userId uuid.UUID, postId uuid.UUID, title string, content string, layer int) error {
	op := "content.service.insertPost"

	logger.StandardDebugF(ctx, op, "Want to insert user:=%v post=%v title=%v content=%v layer=%v",
		userId, postId, title, content, layer)

	err := b.rep.InsertPost(ctx, userId, postId, title, content, layer)
	logger.StandardDebugF(ctx, op, "InsertPost=%v err=%v", postId, err)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (b *Behavior) isUserAuthor(ctx context.Context, userId uuid.UUID) (bool, error) {
	op := "content.service.isUserAuthor"

	logger.StandardDebugF(ctx, op, "Want to check if user %v author", userId)

	role, err := b.rep.GetUserRole(ctx, userId)
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	isAuthor := models.StringToRole(role) == models.Author
	logger.StandardDebugF(ctx, op, "Role = %v userID = %v isAuthor=%v", role, userId, isAuthor)

	if isAuthor {
		return true, nil
	}
	return false, nil
}

func (b *Behavior) checkLayerExist(ctx context.Context, authorId uuid.UUID, layer int) (bool, error) {
	op := "content.service.checkLayerExist"

	if layer == 0 {
		// Посты на нулевом уровне доступны всем
		return true, nil
	}

	// Проверяем есть ли кастомные подписки у этого пользователя на этом уровне
	layerExist, err := b.rep.CheckCustomLayer(ctx, authorId, layer)
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	return layerExist, nil
}
