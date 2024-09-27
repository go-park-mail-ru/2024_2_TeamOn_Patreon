package api

import (
	"fmt"
	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/polina-auth/internal/auth/api/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/polina-auth/internal/auth/api/utils"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/polina-auth/internal/common/logger"
	"net/http"
)

// LoginPost - ручка аутентификации
func LoginPost(w http.ResponseWriter, r *http.Request) {
	op := "auth.api.api_auth.LoginPost"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Парсинг модели вводных данных логина
	var l tModels.Login
	if err := utils.ParseModels(r, &l, op); err != nil {
		// TODO: Дописать отправку модели ошибки с err.msg
		logger.StandardResponse(err.Error(), http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	logger.StandardDebug(fmt.Sprintf("Parsed l={%v}", l), op)

	// Валидация полей вводных данных модели логина
	if _, errV := l.Validate(); errV != nil {
		logger.StandardWarn(
			fmt.Sprintf("Recevied validation error={%v}", errV),
			op)
		// TODO: Дописать отправку модели ошибки с err.msg
		logger.StandardResponse(errV.Error(), http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	logger.StandardDebug(fmt.Sprintf("Validated l={%v}", l), op)

	// достаем behavior из контекста
	b, errM := utils.GetBehaviorCtx(r, op)
	if errM != nil {
		logger.StandardResponse(errM.Error(), http.StatusInternalServerError, r.Host, op)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// достаем инфу по пользователю по username
	// создаем токен пользователя
	tokenString, errM := b.AuthoriseUser(l.Username, l.Password)
	if errM != nil {
		logger.StandardResponse(errM.Error(), http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Сохранение токена в куки
	cookie := utils.CreateCookie(tokenString)

	// Устанавливаем токен в куку
	http.SetCookie(w, &cookie)

	logger.StandardResponse(
		fmt.Sprintf("Successful authorisated user=%v with token='%v'", l.Username, tokenString),
		http.StatusOK,
		r.Host,
		op,
	)

	w.WriteHeader(http.StatusOK)
}
