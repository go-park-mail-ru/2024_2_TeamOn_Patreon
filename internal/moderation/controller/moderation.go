package controller

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/moderation/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

func (h *Handler) ModerationPostComplaintPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// TODO: Достать юзера из кук
	// TODO: Достать пост ид из кук
	// TODO: Валидация пост ид

	// TODO: получить ответ от бизнес-логики

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ModerationPostDecisionPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// TODO: Достать юзера из кук
	// TODO: Достать решение из кук
	// TODO: Валидация решения
	// TODO: получить ответ от бизнес-логики

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ModerationPostGet(w http.ResponseWriter, r *http.Request) {
	op := "moderation.controller.ModerationPostGet"
	ctx := r.Context()

	// TODO: Достать юзера из кук

	// TODO: Достать лимит оффсет и фильтр
	// TODO: Валидация гет параметров

	// TODO: получить ответ от бизнес-логики

	post := models.Post{
		PostID:         uuid.NewV4().String(),
		Title:          "Пост на модерации",
		Content:        "Содержимое поста для модерации",
		AuthorID:       uuid.NewV4().String(),
		AuthorUsername: "автор",
		Status:         "PUBLISHED",
		CreatedAt:      time.Now().String(),
	}

	posts := make([]models.Post, 0)
	posts = append(posts, post)

	utils.SendModel(posts, w, op, ctx)
}
