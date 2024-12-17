/*
 * PushART - Subscription | API
 *
 * API для управления подписками
 */

package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/controller/models/mapper"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func (h *Handler) SubscriptionAuthorIDCustomGet(w http.ResponseWriter, r *http.Request) {
	op := "custom_subscription.controller.SubscriptionAuthorIDCustomGet"

	ctx := r.Context()

	// Получение параметра `authorID` из запроса
	vars := mux.Vars(r)            // Извлекаем параметры из запроса
	authorID := vars[PathAuthorID] // Получаем значение параметра "authorID"

	ok := utils.IsValidUUIDv4(authorID)
	if !ok && authorID != "me" {
		err := errors.Wrap(global.ErrBadRequest, "authorID's invalid")
		logger.StandardDebugF(ctx, op, "get err=%v", err.Error())
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	user, ok := r.Context().Value(global.UserKey).(bModels.User)
	if !ok && authorID == "me" {
		err := errors.Wrap(global.ErrUserNotAuthorized, "user isn't in ctx")
		err = errors.Wrap(err, op)
		logger.StandardDebugF(ctx, op, "get err=%v", err.Error())
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	userID := string(user.UserID)

	if authorID == "me" {
		authorID = userID
	}

	// Достаем кастомные подписки, которые есть у автора
	customSubs, err := h.b.GetCustomSubscription(ctx, authorID)
	if err != nil {
		logger.StandardDebugF(ctx, op, "GetCustomSubscription err=%v", err.Error())
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	var tCustomSubs []*models.CustomSubscription
	tCustomSubs = mapper.MapCommonCustomSubsToTransportSubs(customSubs)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tCustomSubs); err != nil {
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
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	userID := string(user.UserID)

	var acp models.AddCustomSubscription
	if err := utils.ParseModels(r, &acp, op); err != nil {
		err = errors.Wrap(err, "parse model")
		logger.StandardDebugF(ctx, op, "get err=%v", err.Error())
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Валидация полей вводных данных модели
	if err := acp.Validate(); err != nil {
		logger.StandardDebugF(ctx, op, "get err=%v", err.Error())
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	err := h.b.CreateCustomSub(ctx, userID, acp.Title, acp.Description, acp.Layer, acp.Cost)
	if err != nil {
		logger.StandardDebugF(ctx, op, "get err=%v", err.Error())
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// SubscriptionLayersGet обрабатывает запрос на уровни подписок,
// на которых автор может создать новую кастомную подписку
func (h *Handler) SubscriptionLayersGet(w http.ResponseWriter, r *http.Request) {
	op := "custom_subscription.controller.SubscriptionLayersGet"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	ctx := r.Context()

	user, ok := r.Context().Value(global.UserKey).(bModels.User)
	if !ok {
		err := errors.Wrap(global.ErrUserNotAuthorized, "user isn't in ctx")
		err = errors.Wrap(err, op)
		logger.StandardDebugF(ctx, op, "get err=%v", err.Error())
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	userID := string(user.UserID)

	// достаем доступные пользователю подписки
	layers, err := h.b.GetLayerForNewCustomSub(ctx, userID)
	if err != nil {
		logger.StandardDebugF(ctx, op, "get err=%v", err.Error())
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
	}

	var subLayers []*models.SubscriptionLayer
	subLayers = mapper.MapCommonSubLayersToTransportSubLayers(layers)

	w.WriteHeader(http.StatusOK)
	utils.SendModel(subLayers, w, op, ctx)
}
