package controller

import (
	"encoding/json"
	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/models/mapper"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	cModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	utils2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"net/http"
)

func (h *Handler) FeedSubscriptionsGet(w http.ResponseWriter, r *http.Request) {
	op := "internal.content.controller.FeedSubscriptionGet"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	ctx := r.Context()

	// Достаем юзера
	user, ok := r.Context().Value(global.UserKey).(cModels.User)
	if !ok {
		// проставляем http.StatusUnauthorized 401
		logger.StandardResponse("userData not found in context", http.StatusUnauthorized, r.Host, op)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	userId := user.UserID

	// Достаем offset и limit
	// Получение параметров `offset` и `limit` из запроса
	offsetStr := r.URL.Query().Get("offset")
	limitStr := r.URL.Query().Get("limit")

	opt := cModels.NewFeedOpt(offsetStr, limitStr)

	logger.StandardDebugF(op, "Opt=%v", opt)
	// Выполняем бизнес логику
	posts, err := h.b.GetFeedSubscription(ctx, string(userId), opt)
	if err != nil {
		logger.StandardResponse(err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils2.SendStringModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op)
	}

	// мапим посты
	tPosts := mapper.MapCommonPostsToControllerPosts(posts)

	// Отправляем посты
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(tPosts); err != nil {
		logger.StandardResponse(err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
	}
}
