package controller

import (
	"encoding/json"
	"net/http"

	cModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/controller/models"
	global "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
)

// PostAccountUpdate - ручка изменения данных профиля
func (handler *Handler) PostAccountUpdate(w http.ResponseWriter, r *http.Request) {
	op := "internal.account.controller.postAccountUpdate"

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Парсинг данных из json
	newInfo := &cModels.UpdateAccount{}
	if err := json.NewDecoder(r.Body).Decode(&newInfo); err != nil {
		logger.StandardWarnF(op, "json parsing error {%v}", err)
		// Status 400
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Добавить валидацию обновляемых данных!

	// Извлекаем userData из контекста
	userData, ok := r.Context().Value(global.UserKey).(bModels.User)

	if !ok {
		logger.StandardResponse("userData not found in context", http.StatusUnauthorized, r.Host, op)
		// Status 401
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Валидация userID на соответствие стандарту UUIDv4
	if ok := utils.IsValidUUIDv4(string(userData.UserID)); !ok {
		// Status 400
		logger.StandardResponse("invalid userID format", http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	logger.StandardDebugF(op, "Updating userId='%v' username='%v' password='%v' email='%v",
		string(userData.UserID), newInfo.Username, newInfo.Password, newInfo.Email,
	)
	// Обращение к service для записи данных (Может легче было передать сразу всю структуру?)
	if err := handler.serv.PostAccUpdateByID(r.Context(), string(userData.UserID), newInfo.Username, newInfo.Password, newInfo.Email); err != nil {
		logger.StandardWarnF(op, "update data error {%v}", err)
		// Status 500
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Status 200
	w.WriteHeader(http.StatusOK)
}
