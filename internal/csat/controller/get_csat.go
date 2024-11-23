package controller

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
)

func (h *Handler) CsatCheckGet(w http.ResponseWriter, r *http.Request) {
	op := "csat.controller.CsatCheckGet"
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	userIsReady := true
	utils.SendModel(userIsReady, w, op, ctx)
}

func (h *Handler) CsatResultQuestionIDPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
