package controller

import (
	"encoding/json"
	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/models/mapper"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"net/http"
)

func (h *Handler) FeedPopularGet(w http.ResponseWriter, r *http.Request) {
	op := "content.controller.post_post_id_delete.FeedPopularGet"

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Достаем юзера
	user, ok := r.Context().Value(global.UserKey).(models.User)
	if !ok {
		// TODO: Отдать посты для неавторизованного
		return
	}
	userId := string(user.UserID)

	// Получение параметров `offset` и `limit` из запроса
	offsetStr := r.URL.Query().Get("offset")
	limitStr := r.URL.Query().Get("limit")

	opt := models.NewFeedOpt(offsetStr, limitStr)

	posts, err := h.b.GetPopularPostsForUser(userId, opt)
	if err != nil {
		logger.StandardResponse(err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		utils.SendStringModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op)
		return
	}

	popularPosts := mapper.MapCommonPostsToControllerPosts(posts)

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(popularPosts); err != nil {
		logger.StandardResponse(err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
	}
}
