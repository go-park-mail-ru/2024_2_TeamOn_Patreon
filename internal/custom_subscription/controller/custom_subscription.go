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
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
)

func (h *Handler) SubscriptionAuthorIDCustomGet(w http.ResponseWriter, r *http.Request) {
	op := "custom_subscription.controller.SubscriptionAuthorIDCustomGet"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	ctx := r.Context()

	// Получение параметра `authorId` из запроса
	vars := mux.Vars(r)            // Извлекаем параметры из запроса
	authorID := vars[PathAuthorID] // Получаем значение параметра "authorID"

	ok := utils.IsValidUUIDv4(authorID)
	if !ok && authorID != "me" {
		err := errors.Wrap(global.ErrBadRequest, "authorID's invalid")
		utils.SendModel(models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
	}

	var customSubs []*models.CustomSubscription

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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
