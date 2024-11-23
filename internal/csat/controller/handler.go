package controller

import "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/controller/interfaces"

type Handler struct {
	b interfaces.CSATService
}

func New(behavior interfaces.CSATService) *Handler {
	return &Handler{b: behavior}
}

const (
	QueryTimeName = "time"
)
