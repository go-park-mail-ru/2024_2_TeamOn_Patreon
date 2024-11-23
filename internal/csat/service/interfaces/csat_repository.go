package interfaces

import (
	"context"

	repModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/repository/models"
)

type CSATRepository interface {
	// GetQuestion возвращает вопрос для пользователя с его типом
	GetQuestion(ctx context.Context, userID string) (repModels.Question, error)
}
