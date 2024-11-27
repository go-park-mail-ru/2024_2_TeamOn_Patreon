package mapper

import (
	models2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/validate"
)

func MapCommonCustomSubToTransportSub(customSub *models.CustomSubscription) *models2.CustomSubscription {
	title := validate.Sanitize(customSub.Title)
	description := validate.Sanitize(customSub.Description)

	return &models2.CustomSubscription{
		CustomSubscriptionID: customSub.CustomSubscriptionID,
		Title:                title,
		Description:          description,
		Cost:                 customSub.Cost,
		Layer:                customSub.Layer,
	}
}

func MapCommonCustomSubsToTransportSubs(customSubs []*models.CustomSubscription) []*models2.CustomSubscription {
	tCustomSubs := make([]*models2.CustomSubscription, 0, len(customSubs))
	for _, customSub := range customSubs {
		tCustomSubs = append(tCustomSubs, MapCommonCustomSubToTransportSub(customSub))
	}
	return tCustomSubs
}
