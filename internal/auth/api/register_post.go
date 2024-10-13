package api

import (
	"fmt"
	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/api/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/api/utils"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/config"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/logger"
	"net/http"
)

// AuthRegisterPost - ручка регистрации
func (handler *Handler) AuthRegisterPost(w http.ResponseWriter, r *http.Request) {
	op := "auth.api.api_auth.AuthRegisterPost"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Парсинг фронтовой модели данных регистрации
	var p tModels.Reg
	if err := utils.ParseModels(r, &p, op); err != nil {
		// проставляем http.StatusBadRequest
		logger.StandardResponse(err.Error(), config.GetCodeError(err), r.Host, op)
		w.WriteHeader(config.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: config.GetMsgError(err)}, w, op)
		return
	}

	// Валидация полей модели
	if _, err := p.Validate(); err != nil {
		logger.StandardWarnF(op, "Received validation error {%v}", err.Error())
		// проставляем http.StatusBadRequest
		logger.StandardResponse(err.Error(), config.GetCodeError(err), r.Host, op)
		w.WriteHeader(config.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: config.GetMsgError(err)}, w, op)
		return
	}

	// Создание пользователя и генерация токена
	sessionString, err := handler.b.RegisterNewUser(p.Username, p.Password) // передаем username и password
	if err != nil {
		logger.StandardDebugF(op, "Received register error {%v}", err)
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
		fmt.Sprintf("Successful register user=%v with token='%v'", p.Username, sessionString),
		http.StatusOK, r.Host, op)

	w.WriteHeader(http.StatusOK)
}
