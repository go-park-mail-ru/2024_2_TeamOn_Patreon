package controller

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/pkg/validate"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
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

	// Валидация на xss и на слишком длинный
	authorName, err := validate.ValidationAuthorName(authorName)
	if err != nil {
		err = errors.Wrap(err, "authorName's invalid")
		logger.StandardDebugF(ctx, op, "get err=%v", err.Error())
		utils.SendModel(models.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
	}

	// TODO: Поиск авторов бизнес - логике

	// TODO: Валидация на xss и отправка

	w.WriteHeader(http.StatusOK)
}
