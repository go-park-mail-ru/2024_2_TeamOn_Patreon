package controller

import (
	"encoding/json"
	"net/http"

	cModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/controller/models"
	global "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/gorilla/mux"

	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
)

// PostAuthorTip - ручка пожертвований автору
func (handler *Handler) PostAuthorTip(w http.ResponseWriter, r *http.Request) {
	op := "internal.account.controller.PostAuthorUpdateInfo"

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Определяем authorID
	vars := mux.Vars(r)
	authorID := vars["authorID"]

	// Извлекаем userID из контекста
	userData, ok := r.Context().Value(global.UserKey).(bModels.User)
	if !ok {
		logger.StandardResponse("userData not found in context", http.StatusUnauthorized, r.Host, op)
		// Status 401
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	userID := string(userData.UserID)

	// Если пользователь пытается задонатить себе
	if userID == authorID {
		logger.StandardResponse("user can't donate to himself", http.StatusBadRequest, r.Host, op)
		// Status 400
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Валидация userID на соответствие стандарту UUIDv4
	if ok := utils.IsValidUUIDv4(userID); !ok {
		// Status 400
		logger.StandardResponse("invalid userID format", http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Валидация authorID на соответствие стандарту UUIDv4
	if ok := utils.IsValidUUIDv4(authorID); !ok {
		// Status 400
		logger.StandardResponse("invalid authorID format", http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Парсинг данных из json
	tipInfo := &cModels.Tip{}
	if err := json.NewDecoder(r.Body).Decode(&tipInfo); err != nil {
		logger.StandardWarnF(op, "json parsing error {%v}", err)
		// Status 400
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Валидация суммы (не меньше 10 р)
	if tipInfo.Cost < 10 {
		logger.StandardWarnF(op, "the donation amount cannot be less than 10 rubles")
		// Status 400
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Обращение к service
	if err := handler.serv.PostTip(r.Context(), userID, authorID, tipInfo.Cost, tipInfo.Message); err != nil {
		logger.StandardWarnF(op, "update info error {%v}", err)
		// Status 500
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Status 200
	w.WriteHeader(http.StatusOK)
}
