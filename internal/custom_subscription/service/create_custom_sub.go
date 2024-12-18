package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/pkg/validate"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/pkg/errors"
)

func (b *Behavior) CreateCustomSub(ctx context.Context, userID string, title, description string, layer, cost int) error {
	op := "custom_subscription.service.createCustomSub"

	// валидация входных данных
	err := b.validateUserID(userID)
	if err != nil {
		return errors.Wrap(errors.Wrap(err, op), "authorID isn't valid")
	}

	title, description, err = b.validateCustomSub(title, description, layer, cost)
	if err != nil {
		return errors.Wrap(errors.Wrap(err, op), "custom sub isn't valid")
	}

	// Пользователь - автор?
	ok, err := b.isUserAuthor(ctx, userID)
	if err != nil {
		return errors.Wrap(errors.Wrap(err, op), "err - user isn't author")
	}
	if !ok {
		err = global.ErrNotEnoughRights
		return errors.Wrap(errors.Wrap(err, op), "user isn't author")
	}

	// Есть ли кастомная подписка на этом уровне?
	existLayer, err := b.checkCustomLayerExist(ctx, userID, layer)
	if err != nil {
		return errors.Wrap(errors.Wrap(err, op), "layer exists")
	}
	if existLayer {
		return errors.Wrap(errors.Wrap(global.ErrLayerExists, op), "layer exists")
	}

	// Есть ли кастомная подписка на этом уровне?
	existSub, err := b.checkCustomSubTitle(ctx, userID, title)
	if err != nil {
		return errors.Wrap(errors.Wrap(err, op), "title exists")
	}
	if existSub {
		return errors.Wrap(errors.Wrap(global.ErrTitleExists, op), "title exists")
	}

	logger.StandardDebugF(ctx, op,
		"User is Author, Custom Layer doesn't exist, title=%v, description=%v, layer=%v ",
		title, description, layer)

	// создание кастомной подписки
	err = b.createCustomSub(ctx, userID, title, description, layer, cost)
	if err != nil {
		return errors.Wrap(errors.Wrap(err, op), "create custom sub")
	}

	return nil

}

func (b *Behavior) validateUserID(userID string) error {
	op := "custom_subscription.service.validateUserID"
	ok := utils.IsValidUUIDv4(userID)
	if !ok {
		return errors.Wrap(global.ErrBadRequest, op)
	}

	return nil
}

func (b *Behavior) validateCustomSub(title, description string, layer, cost int) (string, string, error) {
	op := "custom_subscription.service.validateCustomSub"
	err := errors.Wrap(global.ErrBadRequest, op)

	title, ok := validate.Title(title)
	if !ok {
		return title, description, errors.Wrap(err, "title")
	}

	description, ok = validate.Description(description)
	if !ok {
		return title, description, errors.Wrap(err, "description")
	}

	ok = validate.Layer(layer)
	if !ok {
		return title, description, errors.Wrap(err, "layer")
	}
	ok = validate.Cost(cost)
	if !ok {
		return title, description, errors.Wrap(err, "cost")
	}
	return title, description, nil
}

func (b *Behavior) isUserAuthor(ctx context.Context, userID string) (bool, error) {
	op := "custom_subscription.service.isUserAuthor"

	role, err := b.rep.GetUserRole(ctx, userID)
	if err != nil {
		return false, errors.Wrap(errors.Wrap(err, op), "get user role")
	}
	if models.StringToRole(role) != models.Author {
		return false, nil
	}

	return true, nil
}

func (b *Behavior) checkCustomLayerExist(ctx context.Context, userID string, layer int) (bool, error) {
	op := "custom_subscription.service.checkCustomLayerExist"

	layers, err := b.rep.GetLayersForAuthor(ctx, userID)
	if err != nil {
		return false, errors.Wrap(errors.Wrap(err, op), "get layers")
	}

	logger.StandardDebugF(ctx, op, "Got layers=%v for author=%v", layers, userID)
	for _, mLayer := range layers {
		if mLayer.Layer == layer {
			return true, nil
		}
	}

	return false, nil
}

func (b *Behavior) createCustomSub(ctx context.Context, userID string, title, description string, layer, cost int) error {
	op := "custom_subscription.service.createCustomSub"

	err := b.rep.CreateCustomSub(ctx, userID, title, description, layer, cost)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

func (b *Behavior) checkCustomSubTitle(ctx context.Context, userID string, title string) (bool, error) {
	op := "custom_subscription.service.checkCustomSubTitle"

	cusSub, err := b.rep.GetCustomSubscriptionByTitle(ctx, userID, title)
	if err != nil {
		return false, errors.Wrap(errors.Wrap(err, op), "get custom sub")
	}
	logger.StandardDebugF(ctx, op, "Got custom sub %v", cusSub)

	return cusSub != nil, nil
}
