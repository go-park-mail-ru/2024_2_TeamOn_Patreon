package controller

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/moderation/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/moderation/controller/models/mapper"
	models2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/moderation/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

func (h *Handler) ModerationPostComplaintPost(w http.ResponseWriter, r *http.Request) {
	op := "moderation.controller.ModerationPostComplaintPost"
	ctx := r.Context()

	// Достать юзера из контекста
	user, ok := r.Context().Value(global.UserKey).(bModels.User)
	if !ok {
		// Не авторизованный пользователь не может жаловаться
		// проставляем http.StatusUnauthorized 401
		logger.StandardResponse(ctx, "userData not found in context", http.StatusUnauthorized, r.Host, op)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Достать пост ид из бади
	var pID models.PostID
	if err := utils.ParseModels(r, &pID, op); err != nil {
		// смысла нет, если пост ид не получен
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Валидация пост ид
	if err := pID.Validate(); err != nil {
		// смысла нет, если пост ид не получен
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Получить ответ от бизнес-логики
	err := h.serv.ComplaintPost(ctx, pID.PostID, string(user.UserID))
	if err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ModerationPostDecisionPost(w http.ResponseWriter, r *http.Request) {
	op := "moderation.controller.ModerationPostDecisionPost"
	ctx := r.Context()

	// Достать юзера из контекста
	user, ok := r.Context().Value(global.UserKey).(bModels.User)
	if !ok {
		// Не авторизованный пользователь не может жаловаться
		// проставляем http.StatusUnauthorized 401
		logger.StandardResponse(ctx, "userData not found in context", http.StatusUnauthorized, r.Host, op)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	userID := string(user.UserID)
	if !utils.IsValidUUIDv4(userID) {
		// Не авторизованный пользователь не может жаловаться
		// проставляем http.StatusUnauthorized 401
		logger.StandardResponse(ctx, "userID in context is incorrect", http.StatusUnauthorized, r.Host, op)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Достать решение из кук
	var dec models.Decision
	if err := utils.ParseModels(r, &dec, op); err != nil {
		// смысла нет, если пост ид не получен
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Валидация решения
	if err := dec.Validate(); err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Получить ответ от бизнес-логики
	err := h.serv.DecisionPost(ctx, dec.PostID, userID, dec.Status)
	if err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ModerationPostGet(w http.ResponseWriter, r *http.Request) {
	op := "moderation.controller.ModerationPostGet"
	ctx := r.Context()

	// Достать юзера из контекста
	user, ok := r.Context().Value(global.UserKey).(bModels.User)
	if !ok {
		// Не авторизованный пользователь не может жаловаться
		// проставляем http.StatusUnauthorized 401
		logger.StandardResponse(ctx, "userData not found in context", http.StatusUnauthorized, r.Host, op)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	userID := string(user.UserID)
	if !utils.IsValidUUIDv4(userID) {
		// Не авторизованный пользователь не может жаловаться
		// проставляем http.StatusUnauthorized 401
		logger.StandardResponse(ctx, "userID in context is incorrect", http.StatusUnauthorized, r.Host, op)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Получение параметров `offset` и `limit` из запроса
	offsetStr := r.URL.Query().Get(QueryOffset)
	limitStr := r.URL.Query().Get(QueryLimit)

	// Получение Фильтра
	filter := r.URL.Query().Get(QueryFilter)
	// Валидация фильтра
	if !models2.CheckStatus(filter) {
		err := errors.Wrap(global.ErrStatusIncorrect, op)
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Валидация гет параметров
	opt := bModels.NewFeedOpt(offsetStr, limitStr)

	// Получить ответ от бизнес-логики
	bPosts, err := h.serv.GetPostsForModeration(ctx, userID, filter, opt.Limit, opt.Offset)
	if err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	posts := mapper.MapBPostsToTPosts(bPosts)

	post := &models.Post{
		PostID:         uuid.NewV4().String(),
		Title:          "Пост на модерации",
		Content:        "Содержимое поста для модерации",
		AuthorID:       uuid.NewV4().String(),
		AuthorUsername: "автор",
		Status:         "PUBLISHED",
		CreatedAt:      time.Now().String(),
	}

	posts = append(posts, post)

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(posts); err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
	}
}
