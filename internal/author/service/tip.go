package service

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/service/models"
	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/service/models"
	global "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (s *Service) CreateTipRequest(ctx context.Context, tipReq models.TipRequest) error {
	op := "internal.author.service.CreateTipRequest"

	// Запрос в repository
	logger.StandardDebugF(ctx, op,
		"want to create subscription request: \tipReqID=%v, \nuserID=%v, \nauthorID=%v, \ncost=%v, \nmessage=%v",
		tipReq.TipReqID, tipReq.UserID, tipReq.AuthorID, tipReq.Cost, tipReq.Message)

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

func (s *Service) RealizeTipRequest(ctx context.Context, tipReqID string, status bool, description string) error {
	op := "internal.author.service.RealizeTipRequest"

	logger.StandardDebugF(ctx, op, "payment status: %v", status)
	// Если не оплачено, подписку не оформляем
	if !status {
		logger.StandardDebugF(ctx, op, "Fail: user did not pay for the order")
		return global.ErrNotPaid
	}

	// Обращение в repository
	logger.StandardDebugF(ctx, op, "want to realize tipReqID request=%v", tipReqID)
	tipReq, err := s.rep.RealizeTipRequest(ctx, tipReqID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful realize tip request=%v", tipReq),
		op)

	// Отправка уведомления автору о новом донате
	if err := s.sendNotificationOfTip(ctx, tipReq.UserID, tipReq.AuthorID, tipReq.Cost, tipReq.Message); err != nil {
		logger.StandardDebugF(ctx, op, "failed send notification to AUTHOR about new tip: %v", err)
	}

	// Отправка уведомления пользователю об отправленном донате
	if err := s.sendNotificationOfTipToDonater(ctx, tipReq.UserID, tipReq.AuthorID, tipReq.Cost); err != nil {
		logger.StandardDebugF(ctx, op, "failed send notification to SUBSCRIBER send tip: %v", err)
	}

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

func (s *Service) sendNotificationOfTipToDonater(ctx context.Context, userID, authorID string, cost int) error {
	op := "internal.author.service.sendNotificationOfTipToDonater"

	authorName, err := s.rep.GetUsername(ctx, authorID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	logger.StandardDebugF(ctx, op, "Want to send notification about successful new tip TO DONATER=%v", userID)

	message := fmt.Sprintf("Пожертвование на сумму %v руб. успешно отправлено автору @%v!", cost, authorName)

	if err := s.rep.SendNotification(ctx, message, userID, userID); err != nil {
		return errors.Wrap(err, op)
	}

	logger.StandardDebugF(ctx, op, "Successful send notification: %v", message)
	return nil
}
