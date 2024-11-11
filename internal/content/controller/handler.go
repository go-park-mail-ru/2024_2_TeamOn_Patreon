package controller

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/interfaces"
)

type Handler struct {
	b interfaces.ContentBehavior
}

func New(behavior interfaces.ContentBehavior) *Handler {
	return &Handler{b: behavior}
}

const (
	// Path param
	authorIDParam string = "authorID"

	// Query param
	offsetParam string = "offset"
	limitParam  string = "limit"
)
