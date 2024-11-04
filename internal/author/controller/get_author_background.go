package controller

import (
	"net/http"

	global "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/gorilla/mux"
)

// GetAuthorBackground - ручка получения фона страницы автора
func (handler *Handler) GetAuthorBackground(w http.ResponseWriter, r *http.Request) {
	op := "internal.account.controller.GetAccountAvatar"

	// Определяем authorID
	vars := mux.Vars(r)
	authorID := vars["authorID"]

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
		logger.StandardResponse("invalid userID format", http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Обращение к service для получения изображения
	avatar, err := handler.serv.GetBackgroundByID(r.Context(), authorID)

	if err != nil {
		logger.StandardDebugF(op, "received background error {%v}", err)
		// Status 500
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Проставляем заголовки
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Disposition", "attachment; filename=\"avatar.jpg\"")

	// Кладём файл в response
	w.Write(avatar)

	// Status 200
	w.WriteHeader(http.StatusOK)
}
