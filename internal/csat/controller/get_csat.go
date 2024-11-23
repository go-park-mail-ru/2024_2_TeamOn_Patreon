package controller

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"net/http"
)

func (h *Handler) CsatCheckGet(w http.ResponseWriter, r *http.Request) {
	op := "csat.controller.CsatCheckGet"
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	userIsReady := true
	utils.SendModel(userIsReady, w, op, ctx)
}

func (h *Handler) CsatQuestionGet(w http.ResponseWriter, r *http.Request) {
	op := "csat.controller.CsatQuestionGet"
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	var question models.Question
	question.QuestionID = "45678"
	question.Question = "Нравится ли вам пользоваться PushArt ?"
	utils.SendModel(question, w, op, ctx)
}

func (h *Handler) CsatResultQuestionIDPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
