package models

import (
	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/service/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/pkg/errors"
)

type SubscriptionRequest struct {
	SubscriptionRequestID string `json:"subscriptionRequestID"`
	AuthorID              string `json:"authorID"`
	MonthCount            int    `json:"monthCount"`
	Layer                 int    `json:"layer"`
}

// MapSubReqToServiceSubReq конвертирует controller SubscriptionRequest в service SubscriptionRequest
func MapSubReqToServiceSubReq(userID string, subReq SubscriptionRequest) sModels.SubscriptionRequest {
	return sModels.SubscriptionRequest{
		SubReqID:   subReq.SubscriptionRequestID,
		UserID:     userID,
		AuthorID:   subReq.AuthorID,
		MonthCount: subReq.MonthCount,
		Layer:      subReq.Layer,
	}
}

func (subReq *SubscriptionRequest) Validate() (bool, error) {
	op := "content.controller.model_subscription_request.Validate"

	if err := subReq.validateAuthorID(); err != nil {
		return false, errors.Wrap(err, op)
	}

	if err := subReq.validateMonthCount(); err != nil {
		return false, errors.Wrap(err, op)
	}

	if err := subReq.validateLayer(); err != nil {
		return false, errors.Wrap(err, op)
	}

	return true, nil
}

func (subReq *SubscriptionRequest) validateAuthorID() error {

	if ok := utils.IsValidUUIDv4(subReq.AuthorID); !ok {
		return global.ErrIsInvalidUUID
	}

	return nil
}

func (subReq *SubscriptionRequest) validateMonthCount() error {

	if subReq.MonthCount < 1 || subReq.MonthCount > 12 {
		return global.ErrInvalidMonthCount
	}

	return nil
}

func (subReq *SubscriptionRequest) validateLayer() error {

	if subReq.Layer < 1 || subReq.Layer > 3 {
		return global.ErrInvalidLayer
	}

	return nil
}
