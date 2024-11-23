package interfaces

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/pkg/models"

	repModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/repository/models"
)

type CSATRepository interface {
	// GetQuestion возвращает вопрос для пользователя с его типом
	GetQuestion(ctx context.Context, userID string) (repModels.Question, error)

	// stat

	GetStatByDays(ctx context.Context, days int) ([]*models.Stat, error)
	GetAllStat(ctx context.Context) ([]*models.Stat, error)
}
