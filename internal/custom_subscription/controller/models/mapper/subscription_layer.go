package mapper

import (
	models2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/validate"
)

func MapCommonSubLayerToTransportSubLayer(subLayer *models.SubscriptionLayer) *models2.SubscriptionLayer {
	layerName := validate.Sanitize(subLayer.LayerName)
	return &models2.SubscriptionLayer{
		Layer:     subLayer.Layer,
		LayerName: layerName,
	}
}

func MapCommonSubLayersToTransportSubLayers(subLayers []*models.SubscriptionLayer) []*models2.SubscriptionLayer {
	tSubLayers := make([]*models2.SubscriptionLayer, 0, len(subLayers))
	for _, subLayer := range subLayers {
		tSubLayers = append(tSubLayers, MapCommonSubLayerToTransportSubLayer(subLayer))
	}
	return tSubLayers
}
