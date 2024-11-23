package controller

import "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/controller/interfaces"

type Handler struct {
	ser interfaces.CSATService
}

func New(ser interfaces.CSATService) *Handler {
	return &Handler{ser: ser}
}
