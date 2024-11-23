package controller

import (
	"net/http"

	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	utils2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
)

func (h *Handler) PostUpdatePost(w http.ResponseWriter, r *http.Request) {
	op := "content.controller.post_post"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ctx := r.Context()

	// Достаем юзера
	user, ok := r.Context().Value(global.UserKey).(bModels.User)
	if !ok {
		// проставляем http.StatusUnauthorized 401
		logger.StandardResponse(ctx, "userData not found in context", http.StatusUnauthorized, r.Host, op)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Парсинг модели вводных данных для добавления поста
	var up models.UpdatePost
	if err := utils2.ParseModels(r, &up, op); err != nil {
		// проставляем http.StatusBadRequest
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils2.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Валидация полей вводных данных модели логина
	if _, err := up.Validate(); err != nil {
		logger.StandardWarnF(ctx, op, "Received validation error={%v}", err)
		// проставляем http.StatusBadRequest
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils2.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	authorID := string(user.UserID)
	// Обновляем пост у соответствующего автора
	err := h.b.UpdatePost(ctx, authorID, up.PostId, up.Title, up.Content)

	if err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils2.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Пост обновлен
	w.WriteHeader(http.StatusCreated)
}
