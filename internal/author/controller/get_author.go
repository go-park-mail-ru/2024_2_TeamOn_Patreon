package controller

import (
	"encoding/json"
	"net/http"

	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller/models"
	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/controller/models"
	global "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	s2Models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"

	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/gorilla/mux"
)

// GetAccount - ручка получения данных профиля
func (handler *Handler) GetAuthor(w http.ResponseWriter, r *http.Request) {
	op := "internal.author.controller.GetAuthor"

	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Определяем authorID
	vars := mux.Vars(r)
	authorID := vars[authorIDParam]
	var userID string

	// Извлекаем userID из контекста
	userData, ok := r.Context().Value(global.UserKey).(s2Models.User)
	if !ok {
		logger.StandardResponse(ctx, "userData not found in context", http.StatusUnauthorized, r.Host, op)
		userID = anon
	} else {
		userID = string(userData.UserID)
	}

	// Если пользователь запрашивает свою страницу
	if authorID == "me" && userID != anon {
		authorID = userID
	}

	// Валидация authorID на соответствие стандарту UUIDv4
	if ok := utils.IsValidUUIDv4(authorID); !ok {
		// Status 400
		logger.StandardResponse(ctx, "invalid authorID format", http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Обращение к service для получения данных автора
	authorData, err := handler.serv.GetAuthorDataByID(r.Context(), authorID)
	if err != nil {
		logger.StandardDebugF(ctx, op, "Received author error {%v}", err)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Обращение к service для получения подписок автора
	subscriptions, err := handler.serv.GetAuthorSubscriptions(r.Context(), authorID)
	if err != nil {
		logger.StandardDebugF(ctx, op, "Received author subscriptions error {%v}", err)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Обращение к service для получения статуса подписки на автора
	isSubscribe, err := handler.serv.GetUserIsSubscribe(r.Context(), authorID, userID)
	if err != nil {
		logger.StandardDebugF(ctx, op, "Received author isSubscribe status error {%v}", err)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	accountPage := sModels.MapAuthorToAuthorPage(authorData, subscriptions, isSubscribe)

	json.NewEncoder(w).Encode(accountPage)
	// Status 200
	w.WriteHeader(http.StatusOK)
}
