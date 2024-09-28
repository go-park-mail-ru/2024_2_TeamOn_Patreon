package api

import (
	"fmt"
	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/polina-auth/internal/auth/api/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/polina-auth/internal/auth/api/utils"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/polina-auth/internal/common/logger"
	"net/http"
)

// AuthRegisterPost - ручка регистрации
func AuthRegisterPost(w http.ResponseWriter, r *http.Request) {
	op := "auth.api.api_auth.AuthRegisterPost"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Парсинг фронтовой модели данных регистрации
	var p tModels.Reg
	if err := utils.ParseModels(r, &p, op); err != nil {
		// проставляем http.StatusBadRequest
		logger.StandardResponse(err.Error(), http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		// отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: "incorrect json"}, w, op)
		return
	}

	// Валидация полей модели
	if _, errV := p.Validate(); errV != nil {
		logger.StandardWarn(fmt.Sprintf("Received validation error {%v}", errV.Error()), op)
		// проставляем http.StatusBadRequest
		logger.StandardResponse(errV.Error(), http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		// отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: errV.GetMsg()}, w, op)
		return
	}

	// Получаем Behavior из контекста
	b, errM := utils.GetBehaviorCtx(r, op)
	if errM != nil {
		logger.StandardWarn(fmt.Sprintf("Received validation error {%v}", errM.Error()), op)
		// проставляем http.StatusBadRequest
		logger.StandardResponse(errM.Error(), http.StatusInternalServerError, r.Host, op)
		w.WriteHeader(http.StatusInternalServerError)
		// отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: errM.GetMsg()}, w, op)
		return
	}

	// Создание пользователя и генерация токена
	tokenString, err := b.RegisterNewUser(p.Username, p.Password) // передаем username и password
	if err != nil {
		logger.StandardDebug(fmt.Sprintf("Received register error {%v}", err), op)
		// проставляем http.StatusBadRequest
		logger.StandardResponse(err.Error(), http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		// отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: err.GetMsg()}, w, op)
		return
	}

	// Сохранение токена в куки
	cookie := utils.CreateCookie(tokenString)

	// Устанавливаем токен в куку
	http.SetCookie(w, &cookie)

	logger.StandardResponse(
		fmt.Sprintf("Successful register user=%v with token='%v'", p.Username, tokenString),
		http.StatusOK, r.Host, op)

	w.WriteHeader(http.StatusOK)
}
