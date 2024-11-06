package controller

import (
	"net/http"

	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/gorilla/mux"
)

func (h Handler) PostsPostIdDelete(w http.ResponseWriter, r *http.Request) {
	op := "content.controller.post_post_id_delete.PostsPostIdDelete"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ctx := r.Context()

	// Достаем postId
	vr := mux.Vars(r)
	postId, ok := vr["postId"]
	logger.StandardDebugF(ctx, op, "postId=%v", postId)
	if !ok {
		// Надо бэд реквест
		err := global.ErrBadRequest
		logger.StandardWarnF(ctx, op, "Received get post from path error={%v} post_id='%v'", err, postId)
		// проставляем http.StatusBadRequest
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendStringModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op)
		return
	}

	// Достаем юзера
	user, ok := r.Context().Value(global.UserKey).(bModels.User)
	if !ok {
		// проставляем http.StatusUnauthorized 401
		logger.StandardResponse(ctx, "userData not found in context", http.StatusUnauthorized, r.Host, op)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Удаляем пост
	userId := user.UserID
	logger.StandardDebugF(ctx, op, "userId=%v postId=%v", userId, postId)
	err := h.b.DeletePost(ctx, string(userId), postId)
	if err != nil {
		logger.StandardWarnF(ctx, op, "Received validation error={%v}", err)
		// проставляем http.StatusBadRequest
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendStringModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
