package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/service/validate"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/pkg/errors"
)

func (b *Behavior) CreatePost(ctx context.Context, authorID string, title string, content string, layer int) (string, error) {
	op := "service.behavior.CreatePost"

	logger.StandardDebugF(ctx, op, "Got authorID:=%v title:=%v content:=%v layer:=%v", authorID, title, content, layer)

	title, content, layer, err := validate.Post(ctx, title, content, layer)
	if err != nil {
		return "", errors.Wrap(err, op)
	}

	newPostID := utils.GenerateUUID()

	if ok := utils.IsValidUUIDv4(authorID); !ok {
		return "", errors.Wrap(global.ErrIsInvalidUUID, op)
	}

	// Проверяем является ли человек автором
	isAuthor, err := b.isUserAuthor(ctx, authorID)
	if err != nil {
		return "", errors.Wrap(err, op)
	}
	if !isAuthor {
		// Почему-то ошибка выделяется красным, хотя все ок
		return "", errors.Wrap(global.ErrNotEnoughRights, op)
	}

	// Проверяем есть ли кастомные подписки у этого пользователя на этом уровне
	layerExist, err := b.checkLayerExist(ctx, authorID, layer)
	if err != nil {
		return "", errors.Wrap(err, op)
	}
	if !layerExist {
		return "", errors.Wrap(global.ErrBadRequest, op)
	}

	// Вставляем пост
	err = b.insertPost(ctx, authorID, newPostID, title, content, layer)
	if err != nil {
		return "", errors.Wrap(err, op)
	}
	return newPostID, nil
}

func (b *Behavior) checkCustomLayer(ctx context.Context, authorID string, layer int) (bool, error) {
	op := "content.service.checkCustomLayer"

	layerExists, err := b.rep.CheckCustomLayer(ctx, authorID, layer)
	if err != nil {
		logger.StandardDebugF(ctx, op, "checkCustomLayer err=%v", err)
		return false, errors.Wrap(err, op)
	}
	return layerExists, nil
}

func (b *Behavior) insertPost(ctx context.Context, userID string, postID string, title string, content string, layer int) error {
	op := "content.service.insertPost"

	logger.StandardDebugF(ctx, op, "Want to insert user:=%v post=%v title=%v content=%v layer=%v",
		userID, postID, title, content, layer)

	err := b.rep.InsertPost(ctx, userID, postID, title, content, layer)
	logger.StandardDebugF(ctx, op, "InsertPost=%v err=%v", postID, err)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (b *Behavior) isUserAuthor(ctx context.Context, userID string) (bool, error) {
	op := "content.service.isUserAuthor"

	logger.StandardDebugF(ctx, op, "Want to check if user %v author", userID)

	role, err := b.rep.GetUserRole(ctx, userID)
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	isAuthor := models.StringToRole(role) == models.Author
	logger.StandardDebugF(ctx, op, "Role = %v userID = %v isAuthor=%v", role, userID, isAuthor)

	if isAuthor {
		return true, nil
	}
	return false, nil
}

func (b *Behavior) checkLayerExist(ctx context.Context, authorID string, layer int) (bool, error) {
	op := "content.service.checkLayerExist"

	if layer == 0 {
		// Посты на нулевом уровне доступны всем
		return true, nil
	}

	// Проверяем есть ли кастомные подписки у этого пользователя на этом уровне
	layerExist, err := b.rep.CheckCustomLayer(ctx, authorID, layer)
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	return layerExist, nil
}
