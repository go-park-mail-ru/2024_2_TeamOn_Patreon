package interfaces

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/pkg/models"

	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/service/models"
)

type CSATService interface {

	// CSATQuestion возвращает запрос для пользователя
	CSATQuestion(ctx context.Context, userID string) (sModels.Question, error)

	// stat

	GetSTATByTime(ctx context.Context, time string) ([]*models.Stat, error)
}
