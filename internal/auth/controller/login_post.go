package controller

import (
	"fmt"
	"net/http"

	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
)

// LoginPost - ручка аутентификации
func (handler *Handler) LoginPost(w http.ResponseWriter, r *http.Request) {
	op := "auth.controller.api_auth.LoginPost"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	ctx := r.Context()

	// Парсинг модели вводных данных логина
	var l tModels.Login
	if err := utils.ParseModels(r, &l, op); err != nil {
		logger.StandardWarnF(ctx, op, "Received parse model error {%v}", err.Error())

		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Валидация полей вводных данных модели логина
	if _, err := l.Validate(); err != nil {
		logger.StandardWarnF(ctx, op, "Received validate error {%v}", err.Error())

		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// достаем инфу по пользователю по username
	// создаем токен пользователя
	tokenString, err := handler.b.LoginUser(ctx, l.Username, l.Password)
	if err != nil {
		logger.StandardWarnF(ctx, op, "Received behavior error {%v}", err.Error())

		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Сохранение токена в куки
	cookie := utils.CreateCookie(tokenString)

	// Устанавливаем токен в куку
	http.SetCookie(w, &cookie)

	logger.StandardResponse(
		ctx,
		fmt.Sprintf("Successful authorisated user=%v with token='%v'", l.Username, tokenString),
		http.StatusOK, r.Host, op)

	w.WriteHeader(http.StatusOK)
}
