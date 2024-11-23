package controller

import (
	"net/http"

	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
)

func (h *Handler) CsatCheckGet(w http.ResponseWriter, r *http.Request) {
	op := "csat.controller.CsatCheckGet"
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Извлекаем userData из контекста
	userData, ok := r.Context().Value(global.UserKey).(bModels.User)

	if !ok {
		logger.StandardResponse(ctx, "userData not found in context", http.StatusUnauthorized, r.Host, op)
		// Status 401
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Валидация userID на соответствие стандарту UUIDv4
	if ok := utils.IsValidUUIDv4(string(userData.UserID)); !ok {
		// Status 400
		logger.StandardResponse(ctx, "invalid userID format", http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Обращение к service
	isCan, err := h.b.IsCanShow(r.Context(), string(userData.UserID))
	if err != nil {
		logger.StandardWarnF(ctx, op, "get canShowSurvey error {%v}", err)
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	var response models.ModelCanShow
	response.IsCanShow = isCan

	utils.SendModel(response, w, op, ctx)
	w.WriteHeader(http.StatusOK)
}
