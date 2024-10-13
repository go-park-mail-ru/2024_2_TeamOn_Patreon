package api

import (
	"fmt"
	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/api/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/api/utils"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/config"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/logger"
	"net/http"
)

// LoginPost - ручка аутентификации
func (handler *Handler) LoginPost(w http.ResponseWriter, r *http.Request) {
	op := "auth.api.api_auth.LoginPost"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Парсинг модели вводных данных логина
	var l tModels.Login
	if err := utils.ParseModels(r, &l, op); err != nil {
		// проставляем http.StatusBadRequest
		logger.StandardResponse(err.Error(), config.GetCodeError(err), r.Host, op)
		w.WriteHeader(config.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: config.GetMsgError(err)}, w, op)
		return
	}

	// Валидация полей вводных данных модели логина
	if _, err := l.Validate(); err != nil {
		logger.StandardWarnF(op, "Received validation error={%v}", err)
		// проставляем http.StatusBadRequest
		logger.StandardResponse(err.Error(), config.GetCodeError(err), r.Host, op)
		w.WriteHeader(config.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: config.GetMsgError(err)}, w, op)
		return
	}

	// достаем инфу по пользователю по username
	// создаем токен пользователя
	// authorise
	// authorize
	sessionString, err := handler.b.AuthoriseUser(l.Username, l.Password)
	if err != nil {
		// проставляем http.StatusBadRequest
		logger.StandardResponse(err.Error(), config.GetCodeError(err), r.Host, op)
		w.WriteHeader(config.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: config.GetMsgError(err)}, w, op)
		return
	}

	// Сохранение токена в куки
	cookie := utils.CreateCookie(sessionString)

	// Устанавливаем токен в куку
	http.SetCookie(w, &cookie)

	logger.StandardResponse(
		fmt.Sprintf("Successful authorisated user=%v with token='%v'", l.Username, sessionString),
		http.StatusOK, r.Host, op)

	w.WriteHeader(http.StatusOK)
}
