package controller

import (
	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	utils2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func (h Handler) AuthorPostAuthorIdGet(w http.ResponseWriter, r *http.Request) {
	op := "content.controller.AuthorPostAuthorIdGet"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Получение параметра `authorId` из запроса
	vars := mux.Vars(r)          // Извлекаем параметры из запроса
	authorId := vars["authorId"] // Получаем значение параметра "authorId"
	if authorId == "" {
		w.WriteHeader(http.StatusBadRequest)
		utils2.SendStringModel(&tModels.ModelError{Message: global.GetMsgError(global.ErrBadRequest)}, w, op)
		return
	}

	// Достаем юзера
	user, ok := r.Context().Value(global.UserKey).(bModels.User)
	if !ok {
		// проставляем http.StatusUnauthorized 401
		//logger.StandardResponse("userData not found in context", http.StatusUnauthorized, r.Host, op)
		if authorId == "me" {
			err := global.ErrBadRequest
			logger.StandardResponse(err.Error(), global.GetCodeError(err), r.Host, op)
			w.WriteHeader(global.GetCodeError(err))
			utils2.SendStringModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op)
			return
		}

		// TODO: достаем список постов для анонима
		// TODO: маппим в транспортные модели
		// TODO: отправляем
	}

	if authorId == "me" {
		// TODO: достаем список постов для конкретного пользователя user

		_ = user
		// TODO: маппим в транспортные модели
		// TODO: отправляем
		return
	}

	// TODO: достаем список постов для самого автора
	// TODO: маппим в транспортные модели
	// TODO: отправляем

	w.WriteHeader(http.StatusOK)
	// сделано на пк
}
