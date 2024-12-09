package controller

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
)

// PostsPostIDCommentsGet - достаем все комменты принадлежащие посту
func (h *Handler) PostsPostIDCommentsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	// TODO: Достаем из контекста пользователя
	// TODO: Достаем и валидируем postID из параметров пути
	// TODO: Достаем и валидируем limit offset из query параметров
	// TODO: Достаем из бизнес логики комменты
	// TODO: Мапим комменты
	// TODO: Отправялем комменты
}

func (h *Handler) PostsPostIDCommentsCreatePost(w http.ResponseWriter, r *http.Request) {
	op := "content.controller.PostsPostIDCommentsCreatePost"
	ctx := r.Context()

	// Достаем из контекста пользователя, если его нет - выходим
	user, ok := r.Context().Value(global.UserKey).(bModels.User)
	if !ok {
		// проставляем http.StatusUnauthorized 401
		logger.StandardResponse(ctx, "userData not found in context", http.StatusUnauthorized, r.Host, op)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Достаем и валидируем postID из параметров пути
	postID, ok := mux.Vars(r)[postIDParam]
	if !utils.IsValidUUIDv4(postID) || !ok {
		err := global.ErrBadRequest
		logger.StandardWarnF(ctx, op, "Received get post from path error={%v} post_id='%v'", err, postID)
		// проставляем http.StatusBadRequest
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Достаем и валидируем UpdateComment из бади
	var ac models.UpdateComment
	if err := utils.ParseModels(r, &ac, op); err != nil {
		// проставляем http.StatusBadRequest
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	if _, err := ac.Validate(); err != nil {
		logger.StandardWarnF(ctx, op, "Received validation error={%v}", err)
		// проставляем http.StatusBadRequest
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Создаем бизнес логикой пост
	commentID, err := h.b.CreateComment(ctx, string(user.UserID), postID, ac.Content)
	if err != nil {
		logger.StandardDebugF(ctx, op, "Error creating comment={%v}", err)
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
	}
	// Отправляем 201
	w.WriteHeader(http.StatusCreated)
	// Отправляем комментИД
	utils.SendModel(models.AddComment{CommentID: commentID}, w, op, ctx)

}

func (h *Handler) PostsCommentsCommentIDUpdatePost(w http.ResponseWriter, r *http.Request) {
	op := "content.controller.PostsPostIDCommentsCreatePost"
	ctx := r.Context()

	// Достаем из контекста пользователя, если его нет - выходим
	user, ok := r.Context().Value(global.UserKey).(bModels.User)
	if !ok {
		// проставляем http.StatusUnauthorized 401
		logger.StandardResponse(ctx, "userData not found in context", http.StatusUnauthorized, r.Host, op)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_ = user
	// TODO: Достаем и валидируем commentID из параметров пути
	// TODO: Достаем и валидируем UpdateComment из бади
	// TODO: Изменяем бизнес логикой пост
	// TODO: Отправляем 200
}

func (h *Handler) PostsCommentsCommentIDDeleteDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// TODO: Достаем из контекста пользователя, если его нет - выходим
	// TODO: Достаем и валидируем  commentID из параметров пути
	// TODO: Удаляем бизнес логикой пост
	// TODO: Отправляем 204
}
