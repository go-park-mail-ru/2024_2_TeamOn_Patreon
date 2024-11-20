package controller

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/pkg/validate"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	models2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
)

func (h *Handler) SearchAuthorNameGet(w http.ResponseWriter, r *http.Request) {
	op := "custom_subscription.controller.SearchAuthorNameGet"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	ctx := r.Context()

	// Получение параметра `authorName` из запроса
	vars := mux.Vars(r)                // Извлекаем параметры из запроса
	authorName := vars[PathAuthorName] // Получаем значение параметра "authorID"

	logger.StandardDebugF(ctx, op, "Received name=%v", authorName)

	// Валидация на xss и на слишком длинный
	authorName, err := validate.ValidationAuthorName(authorName)
	if err != nil {
		err = errors.Wrap(err, "authorName's invalid")
		logger.StandardDebugF(ctx, op, "get err=%v", err.Error())
		utils.SendModel(models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	offset := r.URL.Query().Get(offsetParam)
	limit := r.URL.Query().Get(limitParam)

	opt := models2.NewFeedOpt(offset, limit)

	//  Поиск авторов бизнес - логике
	authorIDs, err := h.b.SearchAuthor(ctx, authorName, opt.Limit, opt.Offset)
	if err != nil {
		err = errors.Wrap(err, "SearchAuthor err")
		logger.StandardDebugF(ctx, op, err.Error())
		utils.SendModel(models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Отправка
	w.WriteHeader(http.StatusOK)
	utils.SendModel(authorIDs, w, op, ctx)
}
