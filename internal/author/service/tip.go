package service

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/service/models"
	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/service/models"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
)

// @@@@@@@@@@@@@@@@@@@@@@@@@@@ Убрать, как только будет готово АПИ оплаты для пожертвований
func (s *Service) PostTip(ctx context.Context, userID, authorID string, cost int, message string) error {
	op := "internal.author.service.PostTip"

	logger.StandardDebugF(ctx, op, "want to save new tip")

	if err := s.rep.NewTip(ctx, userID, authorID, cost, message); err != nil {
		return errors.Wrap(err, op)
	}
	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful send tip: cost=%v, message=%v from user=%v to author=%v", cost, message, userID, authorID),
		op)

	// Отправка уведомления о новом донате
	if err := s.sendNotificationOfTip(ctx, userID, authorID, cost, message); err != nil {
		logger.StandardDebugF(ctx, op, "failed send notification about new subscribe: %v", err)
	}
	return nil
}

func (s *Service) CreateTipRequest(ctx context.Context, tipReq models.TipRequest) error {
	op := "internal.author.service.CreateTipRequest"

	// Запрос в repository
	logger.StandardDebugF(ctx, op, "want to create tip request by user=%v, cost=%v", tipReq.UserID, tipReq.Cost)
	err := s.rep.SaveTipRequest(ctx, sModels.MapServTipReqToRepTipReq(tipReq))
	if err != nil {
		return errors.Wrap(err, op)
	}
	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful create subscription request=%v", tipReq),
		op)

	return nil
}

func (s *Service) sendNotificationOfTip(ctx context.Context, userID, authorID string, cost int, messageInTip string) error {
	op := "internal.author.service.sendNotificationOfTip"

	username, err := s.rep.GetUsername(ctx, userID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	logger.StandardDebugF(ctx, op, "Want to send notification about new tip by user=%v", username)

	var message string

	if messageInTip != "" {
		message = fmt.Sprintf("Пользователь @%v поддержал вас на сумму %v ₽: %v", username, cost, messageInTip)
	} else {
		message = fmt.Sprintf("Пользователь @%v поддержал вас на сумму %v ₽", username, cost)
	}

	if err := s.rep.SendNotification(ctx, message, userID, authorID); err != nil {
		return errors.Wrap(err, op)
	}

	logger.StandardDebugF(ctx, op, "Successful send notification: %v", message)
	return nil
}

func (s *Service) sendNotificationOfTipToDonater(ctx context.Context, userID, authorID string, description string) error {
	op := "internal.author.service.sendNotificationOfTipToDonater"

	authorName, err := s.rep.GetUsername(ctx, authorID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	logger.StandardDebugF(ctx, op, "Want to send notification about successful new tip TO DONATER=%v", userID)

	message := fmt.Sprintf("%v успешно отправлено автору @%v!", description, authorName)

	if err := s.rep.SendNotification(ctx, message, userID, userID); err != nil {
		return errors.Wrap(err, op)
	}

	logger.StandardDebugF(ctx, op, "Successful send notification: %v", message)
	return nil
}
