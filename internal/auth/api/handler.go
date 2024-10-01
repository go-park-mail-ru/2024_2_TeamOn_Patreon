package api

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/behavior/interfaces"
)

type Handler struct {
	b interfaces.AuthBehavior
}

func New(behavior interfaces.AuthBehavior) *Handler {
	return &Handler{b: behavior}
}
