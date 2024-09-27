package api

import (
	"encoding/json"
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
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		logger.StandardWarn(
			fmt.Sprintf("Resived parsing error {%v}", err),
			op,
		)
		// TODO: Дописать отправку модели ошибки с err.msg
		logger.StandardResponse(err.Error(), http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Валидация полей модели
	if _, errV := p.Validate(); errV != nil {
		logger.StandardWarn(
			fmt.Sprintf("Received validation error {%v}", errV.Error()),
			op,
		)
		// TODO: Дописать отправку модели ошибки с err.msg
		logger.StandardResponse(errV.Error(), http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Получаем Behavior из контекста
	b, errM := utils.GetBehaviorCtx(r, op)
	if errM != nil {
		logger.StandardResponse(errM.Error(), http.StatusInternalServerError, r.Host, op)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Создание пользователя и генерация токена
	tokenString, err := b.RegisterNewUser(p.Username, p.Password) // передаем username и password
	if err != nil {
		logger.StandardDebug(
			fmt.Sprintf("Received register error {%v}", err),
			op,
		)
		logger.StandardResponse(err.Error(), http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Сохранение токена в куки
	cookie := utils.CreateCookie(tokenString)

	// Устанавливаем токен в куку
	http.SetCookie(w, &cookie)

	logger.StandardResponse(
		fmt.Sprintf("Successful register user=%v with token='%v'", p.Username, tokenString),
		http.StatusOK,
		r.Host,
		op,
	)

	w.WriteHeader(http.StatusOK)
}
