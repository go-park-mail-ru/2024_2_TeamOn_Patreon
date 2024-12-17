package models

import (
	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/service/models"
)

//go:generate easyjson

// Controller модель пожертвования
//
//easyjson:json
type Tip struct {
	// Сообщение от донатера
	Message string `json:"message,omitempty"`
	// Сумма пожертвования
	Cost int `json:"cost,omitempty"`
}

type TipRequest struct {
	TipReqID string
	UserID   string
	AuthorID string
	Cost     int
	Message  string
}

// MapControllerTipReqToServTipReq конвертирует controller TipRequest в service TipRequest
func MapControllerTipReqToServTipReq(tipReq TipRequest) sModels.TipRequest {
	return sModels.TipRequest{
		TipReqID: tipReq.TipReqID,
		UserID:   tipReq.UserID,
		AuthorID: tipReq.AuthorID,
		Cost:     tipReq.Cost,
		Message:  tipReq.Message,
	}
}
