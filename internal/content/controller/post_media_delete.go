package controller

import (
	"net/http"

	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/gorilla/mux"
)

func (h *Handler) PostDeleteMedia(w http.ResponseWriter, r *http.Request) {
	op := "content.controller.post_delete_media"

	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Достаем юзера
	user, ok := r.Context().Value(global.UserKey).(bModels.User)
	if !ok {
		// проставляем http.StatusUnauthorized 401
		logger.StandardResponse(ctx, "userData not found in context", http.StatusUnauthorized, r.Host, op)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Достаем postID
	vars := mux.Vars(r)
	postID, ok := vars[postIDParam]
	logger.StandardDebugF(ctx, op, "postID=%v", postID)
	if !ok {
		err := global.ErrBadRequest
		logger.StandardWarnF(ctx, op, "Received get post from path error={%v} post_id='%v'", err, postID)
		// проставляем http.StatusBadRequest
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Парсинг модели вводных данных mediaID для добавления поста
	mediaIDs := r.URL.Query()["mediaID"]
	if len(mediaIDs) == 0 {
		err := global.ErrNoFilesToDelete
		logger.StandardWarnF(ctx, op, "Received delete file error={%v} post_id='%v'", err, postID)
		// проставляем http.StatusBadRequest
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Обращение к service
	if err := h.b.DeleteMedia(ctx, string(user.UserID), postID, mediaIDs); err != nil {
		logger.StandardWarnF(ctx, op, "Received delete file error={%v} post_id='%v'", err, postID)
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
	}

	w.WriteHeader(http.StatusOK)
}
