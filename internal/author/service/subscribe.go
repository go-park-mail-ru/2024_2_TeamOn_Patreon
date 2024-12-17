package service

import (
	"context"
	"fmt"
	"strconv"

	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/service/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (s *Service) GetCostSubscription(ctx context.Context, monthCount int, authorID string, layer int) (string, error) {
	op := "internal.author.service.GetCostSubscription"

	logger.StandardDebugF(ctx, op, "want to get cost custom subscription of layer=%v for author=%v", layer, authorID)
	cost, err := s.rep.GetCostCustomSub(ctx, authorID, layer)
	if err != nil {
		return "", errors.Wrap(err, op)
	}

	cost = cost * monthCount

	costStr := strconv.Itoa(cost) + ".00"
	logger.StandardDebugF(ctx, op, "successful get cost = %v", costStr)

	return costStr, nil
}

func (s *Service) CreateSubscriptionRequest(ctx context.Context, subReq sModels.SubscriptionRequest) error {
	op := "internal.author.service.CreateSubscriptionRequest"

	// Автор не может подписаться на самого себя
	logger.StandardDebugF(ctx, op, "want to check author=%v is not user=%v", subReq.AuthorID, subReq.UserID)
	if subReq.AuthorID == subReq.UserID {
		return global.ErrInvalidAuthorID
	}

	// Запрос в repository
	logger.StandardDebugF(ctx, op,
		"want to create subscription request: \nsubReqID=%v, \nuserID=%v, \nauthorID=%v, \nmonth=%v, \nlayer=%v",
		subReq.SubReqID, subReq.UserID, subReq.AuthorID, subReq.MonthCount, subReq.Layer)

	err := s.rep.SaveSubscribeRequest(ctx, sModels.MapServSubReqToRepSubReq(subReq))
	if err != nil {
		return errors.Wrap(err, op)
	}
	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful create subscription request=%v", subReq),
		op)

	return nil
}

func (s *Service) RealizeSubscriptionRequest(ctx context.Context, subReqID string, status bool, description string) error {
	op := "internal.author.service.RealizeSubscriptionRequest"

	logger.StandardDebugF(ctx, op, "payment status: %v", status)
	// Если не оплачено, подписку не оформляем
	if !status {
		logger.StandardDebugF(ctx, op, "Fail: user did not pay for the order")
		return global.ErrNotPaid
	}

	// Обращение в repository
	logger.StandardDebugF(ctx, op, "want to realize subscription request=%v", subReqID)
	subReq, customSubscriptionID, err := s.rep.RealizeSubscribeRequest(ctx, subReqID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful realize subscription request=%v", subReqID),
		op)

	// Отправка уведомления автору о новом подписчике
	if err := s.sendNotificationOfSubscribe(ctx, customSubscriptionID, subReq.UserID); err != nil {
		logger.StandardDebugF(ctx, op, "failed send notification to AUTHOR about new subscribe: %v", err)
	}

	// Отправка уведомления пользователю о оформленной подписке
	if err := s.sendNotificationOfSubscribeToSubscriber(ctx, customSubscriptionID, subReq.UserID, description); err != nil {
		logger.StandardDebugF(ctx, op, "failed send notification to SUBSCRIBER successful save subscription: %v", err)
	}
	return nil
}

func (s *Service) sendNotificationOfSubscribe(ctx context.Context, customSubscriptionID, userID string) error {
	op := "internal.author.service.sendNotificationOfSubscribe"

	username, err := s.rep.GetUsername(ctx, userID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	authorID, customName, err := s.rep.GetCustomSubscriptionInfo(ctx, customSubscriptionID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	message := fmt.Sprintf("Новый подписчик! Пользователь @%v присоединился к вашему сообществу с уровнем: «%v»", username, customName)

	if err := s.rep.SendNotification(ctx, message, userID, authorID); err != nil {
		return errors.Wrap(err, op)
	}

	logger.StandardDebugF(ctx, op, "Successful send notification: %v", message)
	return nil
}

func (s *Service) sendNotificationOfSubscribeToSubscriber(ctx context.Context, customSubscriptionID, userID, description string) error {
	op := "internal.author.service.sendNotificationOfSubscribe"

	authorID, customName, err := s.rep.GetCustomSubscriptionInfo(ctx, customSubscriptionID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	// Имя автора, на которого пользователь подписался
	authorName, err := s.rep.GetUsername(ctx, authorID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	message := fmt.Sprintf("%v на @%v. Уровень подписки: «%v».", description, authorName, customName)

	if err := s.rep.SendNotification(ctx, message, userID, userID); err != nil {
		return errors.Wrap(err, op)
	}

	logger.StandardDebugF(ctx, op, "Successful send notification: %v", message)
	return nil
}
