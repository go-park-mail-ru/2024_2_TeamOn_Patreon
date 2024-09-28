package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/api/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/api/utils"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/logger"
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Получаем Behavior из контекста
	b, errM := utils.GetBehaviorCtx(r, op)
	if errM != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Создание пользователя и генерация токена
	tokenString, err := b.RegisterNewUser(p.Username, p.Password) // передаем username и password
	if err != nil {
		logger.StandardInfo(
			fmt.Sprintf("Received register error {%v}", err),
			op,
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Сохранение токена в куки
	cookie := utils.CreateCookie(tokenString)

	// Устанавливаем токен в куку
	http.SetCookie(w, &cookie)

	logger.StandardInfo(
		fmt.Sprintf("Successful register user=%v with token='%v'", p.Username, tokenString[:15]),
		op,
	)

	w.WriteHeader(http.StatusOK)
}
