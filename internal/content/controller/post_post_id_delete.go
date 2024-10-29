package controller

import (
	"encoding/json"
	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/models/mapper"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"net/http"
	"strconv"
)

func (h Handler) PostsPostIdDelete(w http.ResponseWriter, r *http.Request) {
	op := "content.controller.post_post_id_delete.PostsPostIdDelete"

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Получение параметров `offset` и `limit` из запроса
	offsetStr := r.URL.Query().Get("offset")
	limitStr := r.URL.Query().Get("limit")

	// Установка значений по умолчанию
	offset := 0
	limit := 10

	// Преобразование `offset` и `limit` в целые числа
	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil {
			offset = o
		}
	}
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
		}
	}

	posts, err := h.b.GetPopularPosts(offset, limit)
	if err != nil {
		logger.StandardResponse(err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		utils.SendStringModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op)
		return
	}

	popularPosts := make([]*tModels.Post, 0, len(posts))
	for _, post := range posts {
		popularPosts = append(popularPosts, mapper.MapInterfacePostToTransportPost(post))
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(popularPosts); err != nil {
		logger.StandardResponse(err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
	}
}
