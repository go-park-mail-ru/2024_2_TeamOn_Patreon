package models

type TipRequest struct {
	TipReqID string
	UserID   string
	AuthorID string
	Cost     int
	Message  string
}
