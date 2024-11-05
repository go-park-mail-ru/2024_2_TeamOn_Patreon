package controller

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller/interafces"
)

type Handler struct {
	b interafces.AuthBehavior
}

func New(behavior interafces.AuthBehavior) *Handler {
	return &Handler{b: behavior}
}
