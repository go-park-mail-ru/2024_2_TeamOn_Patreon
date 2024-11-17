/*
 * PushART - Subscription | API
 *
 * API для управления подписками
 */
package controller

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
)

func (h *Handler) SubscriptionAuthorIDCustomGet(w http.ResponseWriter, r *http.Request) {
	op := "custom_subscription.controller.SubscriptionAuthorIDCustomGet"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	ctx := r.Context()

	// Получение параметра `authorID` из запроса
	vars := mux.Vars(r)            // Извлекаем параметры из запроса
	authorID := vars[PathAuthorID] // Получаем значение параметра "authorID"

	ok := utils.IsValidUUIDv4(authorID)
	if !ok && authorID != "me" {
		err := errors.Wrap(global.ErrBadRequest, "authorID's invalid")
		utils.SendModel(models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
	}

	user, ok := r.Context().Value(global.UserKey).(bModels.User)
	if !ok && authorID == "me" {
		err := errors.Wrap(global.ErrUserNotAuthorized, "user isn't in ctx")
		err = errors.Wrap(err, op)
		utils.SendModel(models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
	}

	userID := string(user.UserID)
	_ = userID

	var customSubs []*models.CustomSubscription

	// TODO: Достаем кастомные подписки, которые есть у автора

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(customSubs); err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
	}
}

func (h *Handler) SubscriptionCustomPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) SubscriptionLayersGet(w http.ResponseWriter, r *http.Request) {
	op := "custom_subscription.controller.SubscriptionLayersGet"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	ctx := r.Context()

	user, ok := r.Context().Value(global.UserKey).(bModels.User)
	if !ok {
		err := errors.Wrap(global.ErrUserNotAuthorized, "user isn't in ctx")
		err = errors.Wrap(err, op)
		utils.SendModel(models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
	}

	userID := string(user.UserID)
	_ = userID

	// TODO: достаем доступные пользователю подписки

	var subLayers []*models.SubscriptionLayer

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(subLayers); err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
	}
}
