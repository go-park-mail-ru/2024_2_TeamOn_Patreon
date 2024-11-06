package controller

import (
	"encoding/json"
	"net/http"

	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/models/mapper"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	utils2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/gorilla/mux"
)

func (h Handler) AuthorPostAuthorIdGet(w http.ResponseWriter, r *http.Request) {
	op := "content.controller.AuthorPostAuthorIdGet"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	ctx := r.Context()

	// Получение параметра `authorId` из запроса
	vars := mux.Vars(r)          // Извлекаем параметры из запроса
	authorId := vars["authorId"] // Получаем значение параметра "authorId"
	if authorId == "" {
		w.WriteHeader(http.StatusBadRequest)
		utils2.SendStringModel(&tModels.ModelError{Message: global.GetMsgError(global.ErrBadRequest)}, w, op)
		return
	}

	// Достаем юзера
	userId := ""
	user, ok := r.Context().Value(global.UserKey).(bModels.User)
	if ok {
		userId = string(user.UserID)
	}

	// Достаем offset и limit
	// Получение параметров `offset` и `limit` из запроса
	offsetStr := r.URL.Query().Get("offset")
	limitStr := r.URL.Query().Get("limit")

	opt := bModels.NewFeedOpt(offsetStr, limitStr)

	posts, err := h.b.GetAuthorPosts(ctx, userId, authorId, opt)
	if err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils2.SendStringModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op)
	}

	// мапим посты
	tPosts := mapper.MapCommonPostsToControllerPosts(posts)

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(tPosts); err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
	}
}
