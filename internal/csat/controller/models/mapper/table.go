package mapper

import (
	models2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/pkg/models"
)

func MapStatToTable(stat *models.Stat) *models2.StatTable {
	return &models2.StatTable{
		Theme:  stat.Theme,
		Rating: stat.Rating,
	}
}

func MapStatsToTables(stats []*models.Stat) []*models2.StatTable {
	result := make([]*models2.StatTable, len(stats))
	for i, stat := range stats {
		result[i] = MapStatToTable(stat)
	}
	return result
}
