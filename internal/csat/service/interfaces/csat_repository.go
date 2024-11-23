package interfaces

import (
	"context"

	"time"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/pkg/models"


	repModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/repository/models"
)

type CSATRepository interface {

	// GetQuestion возвращает вопрос для пользователя с его типом
	GetQuestion(ctx context.Context, userID string) (repModels.Question, error)

	// CreateNewAskedQuestion сохраняет новый заданный вопрос
	CreateNewAskedQuestion(ctx context.Context, userID string, question repModels.Question) error

	// SaveRating сохраняет оценку пользователя
	SaveRating(ctx context.Context, userID, questionID string, rating int) error

	// GetTimeLastQuestion возвращает время последнего заданного вопроса
	GetTimeLastQuestion(ctx context.Context, userID string) (time.Time, error)

	// stat

	GetStatByDays(ctx context.Context, days int) ([]*models.Stat, error)
	GetAllStat(ctx context.Context) ([]*models.Stat, error)

}
