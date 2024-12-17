package controller

import (
	"net/http"

	cModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/controller/models"
	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller/models"
	global "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	valid "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/validate"
)

// PostAccountUpdate - ручка изменения данных профиля
func (handler *Handler) PostAccountUpdate(w http.ResponseWriter, r *http.Request) {
	op := "internal.account.controller.postAccountUpdate"

	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Парсинг данных из json
	newInfo := &cModels.UpdateAccount{}
	if err := utils.ParseModels(r, &newInfo, op); err != nil {
		// проставляем http.StatusBadRequest
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Валидация обновляемых данных
	if _, err := newInfo.Validate(); err != nil {
		logger.StandardWarnF(ctx, op, "Received validation error {%v}", err.Error())
		// проставляем http.StatusBadRequest
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, nil)
		return
	}

	// Очищаем только email и username
	newInfo.Email = valid.Sanitize(newInfo.Email)
	newInfo.Username = valid.Sanitize(newInfo.Username)

	// Извлекаем userData из контекста
	userData, ok := r.Context().Value(global.UserKey).(bModels.User)

	if !ok {
		logger.StandardResponse(ctx, "userData not found in context", http.StatusUnauthorized, r.Host, op)
		// Status 401
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Валидация userID на соответствие стандарту UUIDv4
	if ok := utils.IsValidUUIDv4(string(userData.UserID)); !ok {
		// Status 400
		logger.StandardResponse(ctx, "invalid userID format", http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	logger.StandardDebugF(ctx, op, "Updating userId='%v' username='%v' password='%v' email='%v",
		string(userData.UserID), newInfo.Username, newInfo.Password, newInfo.Email,
	)

	if newInfo.Password != "" {
		if err := handler.serv.UpdatePassword(r.Context(), string(userData.UserID), newInfo.OldPassword, newInfo.Password); err != nil {
			logger.StandardWarnF(ctx, op, "update Password error {%v}", err)
			w.WriteHeader(global.GetCodeError(err))
			utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
			return
		}
	}

	if newInfo.Username != "" {
		if err := handler.serv.UpdateUsername(r.Context(), string(userData.UserID), newInfo.Username); err != nil {
			logger.StandardWarnF(ctx, op, "update Username error {%v}", err)
			w.WriteHeader(global.GetCodeError(err))
			utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
			return
		}
	}

	if newInfo.Email != "" {
		if err := handler.serv.UpdateEmail(r.Context(), string(userData.UserID), newInfo.Email); err != nil {
			logger.StandardWarnF(ctx, op, "update Email error {%v}", err)
			w.WriteHeader(global.GetCodeError(err))
			utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
			return
		}
	}
	// Status 200
	w.WriteHeader(http.StatusOK)
}
