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
		logger.StandardDebugF(ctx, op, "get err=%v", err.Error())
		utils.SendModel(models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	user, ok := r.Context().Value(global.UserKey).(bModels.User)
	if !ok && authorID == "me" {
		err := errors.Wrap(global.ErrUserNotAuthorized, "user isn't in ctx")
		err = errors.Wrap(err, op)
		logger.StandardDebugF(ctx, op, "get err=%v", err.Error())
		utils.SendModel(models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
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
	op := "custom_subscription.controller.SubscriptionCustomPost"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	ctx := r.Context()

	user, ok := r.Context().Value(global.UserKey).(bModels.User)
	if !ok {
		err := errors.Wrap(global.ErrUserNotAuthorized, "user isn't in ctx")
		err = errors.Wrap(err, op)
		logger.StandardDebugF(ctx, op, "get err=%v", err.Error())
		utils.SendModel(models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	userID := string(user.UserID)

	var acp models.AddCustomSubscription
	if err := utils.ParseModels(r, &acp, op); err != nil {
		err = errors.Wrap(err, "parse model")
		logger.StandardDebugF(ctx, op, "get err=%v", err.Error())
		utils.SendModel(models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Валидация полей вводных данных модели
	if err := acp.Validate(); err != nil {
		logger.StandardDebugF(ctx, op, "get err=%v", err.Error())
		utils.SendModel(models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	err := h.b.CreateCustomSub(ctx, userID, acp.Title, acp.Description, acp.Layer, acp.Cost)
	if err != nil {
		logger.StandardDebugF(ctx, op, "get err=%v", err.Error())
		utils.SendModel(models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) SubscriptionLayersGet(w http.ResponseWriter, r *http.Request) {
	op := "custom_subscription.controller.SubscriptionLayersGet"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	ctx := r.Context()

	user, ok := r.Context().Value(global.UserKey).(bModels.User)
	if !ok {
		err := errors.Wrap(global.ErrUserNotAuthorized, "user isn't in ctx")
		err = errors.Wrap(err, op)
		logger.StandardDebugF(ctx, op, "get err=%v", err.Error())
		utils.SendModel(models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
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
