package controller

import (
	"encoding/json"
	"net/http"

	global "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/gorilla/mux"
)

// GetAccount - ручка получения данных профиля
func (handler *Handler) GetAuthor(w http.ResponseWriter, r *http.Request) {
	op := "internal.author.controller.GetAuthor"

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Определяем authorID
	vars := mux.Vars(r)
	authorID := vars["authorID"]

	// Если пользователь запрашивает свою страницу
	if authorID == "me" {
		// Извлекаем userID из контекста
		userData, ok := r.Context().Value(global.UserKey).(bModels.User)
		if !ok {
			logger.StandardResponse("userData not found in context", http.StatusUnauthorized, r.Host, op)
			// Status 401
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		authorID = string(userData.UserID)
	}

	// Валидация authorID на соответствие стандарту UUIDv4
	if ok := utils.IsValidUUIDv4(authorID); !ok {
		// Status 400
		logger.StandardResponse("invalid authorID format", http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Обращение к service для получения данных
	authorData, err := handler.serv.GetAuthorDataByID(r.Context(), authorID)
	if err != nil {
		logger.StandardDebugF(op, "Received author error {%v}", err)
		// Status 500
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(authorData)
	// Status 200
	w.WriteHeader(http.StatusOK)
}
