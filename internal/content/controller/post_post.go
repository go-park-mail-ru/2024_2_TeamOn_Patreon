package controller

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/models"
	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	utils2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"net/http"
)

// PostPost - Позволяет создать новый пост, добавить заголовок, описание, контент
func (h Handler) PostPost(w http.ResponseWriter, r *http.Request) {
	op := "content.controller.post_post"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Достаем юзера
	user, ok := r.Context().Value(global.UserKey).(bModels.User)
	if !ok {
		// проставляем http.StatusUnauthorized 401
		logger.StandardResponse("userData not found in context", http.StatusUnauthorized, r.Host, op)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Парсинг модели вводных данных для добавления поста
	var ap models.AddPost
	if err := utils2.ParseModels(r, &ap, op); err != nil {
		// проставляем http.StatusBadRequest
		logger.StandardResponse(err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils2.SendStringModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op)
		return
	}

	// Валидация полей вводных данных модели логина
	if _, err := ap.Validate(); err != nil {
		logger.StandardWarnF(op, "Received validation error={%v}", err)
		// проставляем http.StatusBadRequest
		logger.StandardResponse(err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils2.SendStringModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op)
		return
	}

	authorId := string(user.UserID)
	postID, err := h.b.CreatePost(authorId, ap.Title, ap.Content, ap.Layer)
	if err != nil {
		logger.StandardResponse(err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		utils2.SendStringModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op)
		return
	}

	w.WriteHeader(http.StatusCreated)
	utils2.SendStringModel(&tModels.PostId{PostId: postID}, w, op)
}