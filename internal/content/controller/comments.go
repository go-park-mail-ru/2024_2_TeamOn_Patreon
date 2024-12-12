package controller

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/models/mapper"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
)

// PostsPostIDCommentsGet - достаем все комменты принадлежащие посту
func (h *Handler) PostsPostIDCommentsGet(w http.ResponseWriter, r *http.Request) {
	op := "content.controller.PostsPostIDCommentsGet"
	ctx := r.Context()

	// Достаем из контекста пользователя
	var userID string
	user, ok := r.Context().Value(global.UserKey).(bModels.User)
	if ok {
		userID = string(user.UserID)
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

	// Получение параметров `offset` и `limit` из запроса
	offsetStr := r.URL.Query().Get(offsetParam)
	limitStr := r.URL.Query().Get(limitParam)

	opt := bModels.NewFeedOpt(offsetStr, limitStr)

	// Достаем из бизнес логики комменты
	comments, err := h.b.GetComments(ctx, userID, postID, opt)
	if err != nil {
		logger.StandardWarnF(ctx, op, "Get comments error={%v}", err)
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}
	// Мапим комменты
	tComments := mapper.MapCommonCommentsToControllerComments(comments)

	// Отправялем комменты
	utils.SendModel(tComments, w, op, ctx)
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

	// Достаем и валидируем commentID из параметров пути
	commentID, ok := mux.Vars(r)[commentIDParam]
	if !utils.IsValidUUIDv4(commentID) || !ok {
		err := global.ErrBadRequest
		logger.StandardWarnF(ctx, op, "Received get post from path error={%v} post_id='%v'", err, commentID)
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

	// Изменяем бизнес логикой пост
	err := h.b.UpdateComment(ctx, string(user.UserID), commentID, ac.Content)
	if err != nil {
		logger.StandardDebugF(ctx, op, "Error updating comment={%v}", err)
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Отправляем 200
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) PostsCommentsCommentIDDeleteDelete(w http.ResponseWriter, r *http.Request) {
	op := "content.controller.PostsCommentsCommentIDDeleteDelete"
	ctx := r.Context()

	// Достаем из контекста пользователя, если его нет - выходим
	user, ok := r.Context().Value(global.UserKey).(bModels.User)
	if !ok {
		// проставляем http.StatusUnauthorized 401
		logger.StandardResponse(ctx, "userData not found in context", http.StatusUnauthorized, r.Host, op)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Достаем и валидируем  commentID из параметров пути
	commentID, ok := mux.Vars(r)[commentIDParam]
	if !utils.IsValidUUIDv4(commentID) || !ok {
		err := global.ErrBadRequest
		logger.StandardWarnF(ctx, op, "Received get post from path error={%v} post_id='%v'", err, commentID)
		// проставляем http.StatusBadRequest
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Удаляем бизнес логикой пост
	err := h.b.DeleteComment(ctx, string(user.UserID), commentID)
	if err != nil {
		logger.StandardDebugF(ctx, op, "Error deleting comment={%v}", err)
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Отправляем 204
	w.WriteHeader(http.StatusNoContent)
}
