package controller

import (
	"encoding/json"
	"net/http"

	global "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
)

// HandlerGetAccount - ручка получения данных профиля
func (handler *Handler) GetAccount(w http.ResponseWriter, r *http.Request) {
	op := "account.api.api_account"

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Извлекаем userData из контекста
	userData, ok := r.Context().Value(global.UserKey).(bModels.User)

	if !ok {
		// проставляем http.StatusUnauthorized 401
		logger.StandardResponse("userData not found in context", http.StatusUnauthorized, r.Host, op)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Валидация userID на соответствие стандарту UUIDv4
	if !utils.IsValidUUIDv4(string(userData.UserID)) {
		// проставляем http.StatusBadRequest 400
		logger.StandardResponse("invalid userID format", http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Обращение к service для получения данных
	accountData, err := handler.serv.GetAccDataByID(r.Context(), string(userData.UserID))
	if err != nil {
		logger.StandardDebugF(op, "Received account error {%v}", err)
		// проставляем http.StatusInternalServerError 500
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(accountData)
	w.WriteHeader(http.StatusOK)
}
