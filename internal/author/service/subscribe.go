package service

import (
	"context"
	"fmt"

	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/service/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (s *Service) CreateSubscriptionRequest(ctx context.Context, subReq sModels.SubscriptionRequest) (string, error) {
	op := "internal.author.service.CreateSubscriptionRequest"

	// Автор не может подписаться на самого себя
	logger.StandardDebugF(ctx, op, "want to check author=%v is not user=%v", subReq.AuthorID, subReq.UserID)
	if subReq.AuthorID == subReq.UserID {
		return "", global.ErrInvalidAuthorID
	}

	// Запрос в repository
	logger.StandardDebugF(ctx, op, "want to create subscription request by user=%v, month=%v, layer=%v", subReq.UserID, subReq.MonthCount, subReq.Layer)
	subReqID, err := s.rep.CreateSubscribeRequest(ctx, sModels.MapServSubReqToRepSubReq(subReq))
	if err != nil {
		return "", errors.Wrap(err, op)
	}
	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful create subscription request=%v", subReq),
		op)

	return subReqID, nil
}

func (s *Service) RealizeSubscriptionRequest(ctx context.Context, subReqID string) error {
	op := "internal.author.service.RealizeSubscriptionRequest"

	// Обращение в repository
	logger.StandardDebugF(ctx, op, "want to realize subscription request=%v", subReqID)
	if err := s.rep.RealizeSubscribeRequest(ctx, subReqID); err != nil {
		return errors.Wrap(err, op)
	}

	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful realize subscription request=%v", subReqID),
		op)

	return nil
}
