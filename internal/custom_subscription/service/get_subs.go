package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/pkg/errors"
)

func (b *Behavior) GetLayerForNewCustomSub(ctx context.Context, userID string) ([]*models.SubscriptionLayer, error) {
	op := "custom_subscription.service.GetLayerForNewCustomSub"
	layers, err := b.rep.GetFreeLayersForAuthor(ctx, userID)

	newLayers := make([]*models.SubscriptionLayer, 0)
	for _, layer := range layers {
		if layer.Layer != 0 {
			newLayers = append(newLayers, layer)
		}
	}
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	return newLayers, nil
}

func (b *Behavior) GetCustomSubscription(ctx context.Context, authorID string) ([]*models.CustomSubscription, error) {
	op := "custom_subscription.service.GetCustomSubscription"

	ok := utils.IsValidUUIDv4(authorID)
	if !ok {
		return nil, errors.Wrap(global.ErrIsInvalidUUID, op)
	}

	customSubs, err := b.rep.GetCustomSubscription(ctx, authorID)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	return customSubs, nil
}
