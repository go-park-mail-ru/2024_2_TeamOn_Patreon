package controller

import (
	"fmt"
	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	utils2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"net/http"
)

// LoginPost - ручка аутентификации
func (handler *Handler) LoginPost(w http.ResponseWriter, r *http.Request) {
	op := "auth.controller.api_auth.LoginPost"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Парсинг модели вводных данных логина
	var l tModels.Login
	if err := utils2.ParseModels(r, &l, op); err != nil {
		// проставляем http.StatusBadRequest
		logger.StandardResponse(err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils2.SendStringModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op)
		return
	}

	// Валидация полей вводных данных модели логина
	if _, err := l.Validate(); err != nil {
		logger.StandardWarnF(op, "Received validation error={%v}", err)
		// проставляем http.StatusBadRequest
		logger.StandardResponse(err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils2.SendStringModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op)
		return
	}

	// достаем инфу по пользователю по username
	// создаем токен пользователя
	// authorise
	// authorize
	tokenString, err := handler.b.AuthoriseUser(l.Username, l.Password)
	if err != nil {
		// проставляем http.StatusBadRequest
		logger.StandardResponse(err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils2.SendStringModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op)
		return
	}

	// Сохранение токена в куки
	cookie := utils2.CreateCookie(tokenString)

	// Устанавливаем токен в куку
	http.SetCookie(w, &cookie)

	logger.StandardResponse(
		fmt.Sprintf("Successful authorisated user=%v with token='%v'", l.Username, tokenString),
		http.StatusOK, r.Host, op)

	w.WriteHeader(http.StatusOK)
}
