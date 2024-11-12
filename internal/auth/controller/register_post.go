package controller

import (
	"fmt"
	"net/http"

	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
)

// AuthRegisterPost - ручка регистрации
func (handler *Handler) AuthRegisterPost(w http.ResponseWriter, r *http.Request) {
	op := "auth.controller.api_auth.AuthRegisterPost"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	ctx := r.Context()

	// Парсинг фронтовой модели данных регистрации
	var p tModels.Reg
	if err := utils.ParseModels(r, &p, op); err != nil {
		// проставляем http код ошибки
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Валидация полей модели
	if _, err := p.Validate(); err != nil {
		logger.StandardWarnF(ctx, op, "Received validation error {%v}", err.Error())
		// проставляем http.StatusBadRequest
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Создание пользователя и генерация токена
	tokenString, err := handler.b.RegisterNewUser(ctx, p.Username, p.Password) // передаем username и password
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
		fmt.Sprintf("Successful register user=%v with token='%v'", p.Username, tokenString),
		http.StatusOK, r.Host, op)

	w.WriteHeader(http.StatusOK)
}
