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
	logger.StandardDebugF(ctx, op, "want to create subscription request by user=%v, month=%v, layer=%v", subReq.UserID, subReq.MonthCount, subReq.Layer)
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

func (s *Service) RealizeSubscriptionRequest(ctx context.Context, subReqID, userID string) error {
	op := "internal.author.service.RealizeSubscriptionRequest"

	// Обращение в repository
	logger.StandardDebugF(ctx, op, "want to realize subscription request=%v", subReqID)
	customSubscriptionID, err := s.rep.RealizeSubscribeRequest(ctx, subReqID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful realize subscription request=%v", subReqID),
		op)

	// Отправка уведомления о новой подписке
	if err := s.sendNotificationOfSubscribe(ctx, customSubscriptionID, userID); err != nil {
		logger.StandardDebugF(ctx, op, "failed send notification about new subscribe: %v", err)
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
