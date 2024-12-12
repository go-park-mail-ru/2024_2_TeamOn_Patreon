package models

import repModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/repository/models"

type TipRequest struct {
	TipReqID string
	UserID   string
	AuthorID string
	Cost     int
	Message  string
}

// MapServTipReqToRepTipReq конвертирует service TipRequest в repository TipRequest
func MapServTipReqToRepTipReq(tipReq TipRequest) repModels.TipRequest {
	return repModels.TipRequest{
		TipReqID: tipReq.TipReqID,
		UserID:   tipReq.UserID,
		AuthorID: tipReq.AuthorID,
		Cost:     tipReq.Cost,
		Message:  tipReq.Message,
	}
}
