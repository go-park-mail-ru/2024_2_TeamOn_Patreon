package mapper

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/models"
	models2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
)

func MapCommonMediaToTransportMedia(post models2.Media) *models.Media {
	return &models.Media{
		MediaID:   post.MediaID,
		MediaType: post.MediaType,
		MediaURL:  post.MediaURL,
	}
}

func MapCommonMediaSToControllerMedias(medias []*models2.Media) []*models.Media {
	tMedias := make([]*models.Media, 0, len(medias))
	for _, media := range medias {
		tMedias = append(tMedias, MapCommonMediaToTransportMedia(*media))
	}
	return tMedias
}
