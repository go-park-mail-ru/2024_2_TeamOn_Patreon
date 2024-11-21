package models

import (
	repModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/repository/models"
)

type SubscriptionRequest struct {
	UserID     string
	AuthorID   string
	MonthCount int
	Layer      int
}

// MapServSubReqToRepSubReq конвертирует service SubscriptionRequest в repository SubscriptionRequest
func MapServSubReqToRepSubReq(subReq SubscriptionRequest) repModels.SubscriptionRequest {
	return repModels.SubscriptionRequest{
		UserID:     subReq.UserID,
		AuthorID:   subReq.AuthorID,
		MonthCount: subReq.MonthCount,
		Layer:      subReq.Layer,
	}
}
