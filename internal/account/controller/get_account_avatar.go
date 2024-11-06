package controller

import (
	"net/http"

	global "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/gorilla/mux"
)

// GetAccountAvatar - ручка получения аватарки пользователя
func (handler *Handler) GetAccountAvatar(w http.ResponseWriter, r *http.Request) {
	op := "internal.account.controller.GetAccountAvatar"

	ctx := r.Context()
	// Определяем userID
	vars := mux.Vars(r)
	userID := vars["userID"]

	if userID == "me" {
		// Извлекаем userID из контекста
		userData, ok := ctx.Value(global.UserKey).(bModels.User)
		if !ok {
			logger.StandardResponse(ctx, "userData not found in context", http.StatusUnauthorized, r.Host, op)
			// Status 401
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		userID = string(userData.UserID)
	}

	// Валидация userID на соответствие стандарту UUIDv4
	if ok := utils.IsValidUUIDv4(userID); !ok {
		// Status 400
		logger.StandardResponse(ctx, "invalid userID format", http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Обращение к service для получения изображения
	avatar, err := handler.serv.GetAvatarByID(r.Context(), userID)

	if err != nil {
		logger.StandardDebugF(ctx, op, "received avatar error {%v}", err)
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
