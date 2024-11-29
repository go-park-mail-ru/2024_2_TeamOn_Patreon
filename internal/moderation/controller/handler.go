package controller

import "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/moderation/controller/interfaces"

type Handler struct {
	serv interfaces.ModerationService
}

func New(serv interfaces.ModerationService) *Handler {
	return &Handler{
		serv: serv,
	}
}
