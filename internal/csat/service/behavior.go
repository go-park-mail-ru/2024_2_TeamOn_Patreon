package service

import (
	"context"
	"fmt"
	"time"

	rInterfaces "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/service/interfaces"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
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

const (
	Infinity = "infinity"
)

func (b *Behavior) CSATQuestion(ctx context.Context, userID string) (sModels.Question, error) {
	op := "internal.csat.behavior.CSATQuestion"

	// получаем данные пользователя из rep
	logger.StandardDebugF(ctx, op, "want to get question for userID = %v", userID)

	// обращение в репозиторий для получения вопроса
	question, err := b.rep.GetQuestion(ctx, userID)
	if err != nil {
		return sModels.Question{}, errors.Wrap(err, op)
	}

	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful get question=(%v) for userID=%v", question.Question, userID),
		op)

	// Если вопрос успешно получен, то создаём запись в таблице заданных вопросов
	logger.StandardDebugF(ctx, op, "want to save asked question for userID = %v", userID)
	if err := b.rep.CreateNewAskedQuestion(ctx, userID, question); err != nil {
		return sModels.Question{}, errors.Wrap(err, op)
	}

	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful save asked question=(%v) for userID=%v", question.Question, userID),
		op)

	return sModels.MapRepQuestionToServQuestion(question), nil
}

func (b *Behavior) SaveRating(ctx context.Context, userID, questionID string, rating int) error {
	op := "internal.csat.behavior.SaveRating"

	// получаем данные пользователя из rep
	logger.StandardDebugF(ctx, op, "want to save new rating question for userID = %v", userID)

	// Валидация
	if rating < 0 || rating > 5 {
		return global.ErrInvalidRating
	}

	// обращение в репозиторий
	if err := b.rep.SaveRating(ctx, userID, questionID, rating); err != nil {
		return errors.Wrap(err, op)
	}

	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful save rating=%v for questionID=(%v) for userID=%v", rating, questionID, userID),
		op)

	return nil
}

func (b *Behavior) IsCanShow(ctx context.Context, userID string) (bool, error) {
	op := "internal.csat.behavior.IsCanShow"

	// Константа, с каким периодом задаём вопросы
	periodAskQuestions := 5 * time.Minute

	// получаем данные пользователя из rep
	logger.StandardDebugF(ctx, op, "want to get IsCanShow question for user=%v", userID)

	lastTime, err := b.rep.GetTimeLastQuestion(ctx, userID)
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful get last time question=%v", lastTime),
		op)

	logger.StandardDebugF(ctx, op, "want to compare time for user=%v", userID)

	// Проверяем разницу по времени (что прошло period минут)

	if err != nil {
		fmt.Println("Ошибка при парсинге времени:", err)
		return false, errors.Wrap(err, op)
	}

	// Получение текущего времени
	currentTime := time.Now()

	// Вычисление разницы во времени
	duration := currentTime.Sub(lastTime)

	var isCanShow bool
	// Проверка, прошло ли 10 минут
	if duration > periodAskQuestions {
		isCanShow = true
	} else {
		isCanShow = false
	}

	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful get IsCanShow=%v", isCanShow),
		op)

	return isCanShow, nil
}
