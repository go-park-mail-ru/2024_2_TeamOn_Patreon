package service

import (
	"context"
	"fmt"

	rInterfaces "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/service/interfaces"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"

	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/service/models"
)

type Behavior struct {
	rep rInterfaces.CSATRepository
}

func New(repository rInterfaces.CSATRepository) *Behavior {
	return &Behavior{repository}
}

func (b *Behavior) CSATQuestion(ctx context.Context, userID string) (sModels.Question, error) {
	op := "internal.auth.behavior.CSATQuestion"

	// получаем данные пользователя из rep
	logger.StandardDebugF(ctx, op, "want to get question for userID = %v", userID)

	// обращение в репозиторий
	question, err := b.rep.GetQuestion(ctx, userID)
	if err != nil {
		return sModels.Question{}, errors.Wrap(err, op)
	}

	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful get question=(%v) for userID=%v", question.Question, userID),
		op)

	return sModels.MapRepQuestionToServQuestion(question), nil
}
