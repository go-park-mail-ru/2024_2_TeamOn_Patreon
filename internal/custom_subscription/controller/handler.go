package controller

import "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/controller/interfaces"

type Handler struct {
	b interfaces.CustomSubscriptionService
}

func New(behavior interfaces.CustomSubscriptionService) *Handler {
	return &Handler{b: behavior}
}

const (
	PathAuthorID = "authorID"
)
