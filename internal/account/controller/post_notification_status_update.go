package controller

import (
	"encoding/json"
	"net/http"

	models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/controller/models"
	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller/models"
	global "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	pkgModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
)

// GetNotifications - ручка изменения статуса уведомления на "прочитано"
func (handler *Handler) PostNotificationStatusUpdate(w http.ResponseWriter, r *http.Request) {
	op := "account.controller.PostNotificationStatusUpdate"
	ctx := r.Context()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Извлекаем userData из контекста
	userData, ok := ctx.Value(global.UserKey).(pkgModels.User)
	if !ok {
		err := global.ErrUserNotAuthorized
		logger.StandardResponse(ctx, "Auth failed: fail get user from ctx", global.GetCodeError(err), r.Host, op)
		// Status 401
		w.WriteHeader(global.GetCodeError(err))
		// Отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}
	userID := string(userData.UserID)

	// Валидация userID на соответствие стандарту UUIDv4
	if ok := utils.IsValidUUIDv4(string(userData.UserID)); !ok {
		err := global.ErrIsInvalidUUID
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		// Status 400
		w.WriteHeader(global.GetCodeError(err))
		// Отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Парсинг данных из json
	notification := &models.NotificationID{}
	if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
		logger.StandardWarnF(ctx, op, "json parsing error {%v}", err)
		// Status 400
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	notificationID := notification.ID

	// Валидация обновляемых данных
	if ok := utils.IsValidUUIDv4(notificationID); !ok {
		err := global.ErrIsInvalidUUID
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		// Status 400
		w.WriteHeader(global.GetCodeError(err))
		// Отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	err := handler.serv.ReadNotification(ctx, userID, notificationID)
	if err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	w.WriteHeader(http.StatusOK)
}
