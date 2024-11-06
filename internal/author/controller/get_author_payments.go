package controller

import (
	"encoding/json"
	"net/http"

	cModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/controller/models"
	global "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"

	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
)

// GetAccount - ручка получения данных профиля
func (handler *Handler) GetAuthorPayments(w http.ResponseWriter, r *http.Request) {
	op := "internal.author.controller.GetAuthorPayments"

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

	// Валидация authorID на соответствие стандарту UUIDv4
	if ok := utils.IsValidUUIDv4(string(userData.UserID)); !ok {
		// Status 400
		logger.StandardResponse(ctx, "invalid authorID format", http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Обращение к service для получения данных
	amountPayments, err := handler.serv.GetAuthorPayments(r.Context(), string(userData.UserID))
	if err != nil {
		logger.StandardDebugF(ctx, op, "Received author payments error {%v}", err)
		// Status 500
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	payments := cModels.Payments{
		Amount: amountPayments,
	}
	json.NewEncoder(w).Encode(payments)
	// Status 200
	w.WriteHeader(http.StatusOK)
}
