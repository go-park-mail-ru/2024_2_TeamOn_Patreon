package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/controller/models"
	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	pkgModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
)

// GetNotifications - ручка получения уведомлений пользователя
// status - статус уведомлений. Если NOTREAD - получить непрочитанные. Если не передаётся - получить все
// offsetParam - имя для параметра ограничения, limitParam - имя для параметра смещения

func (handler *Handler) GetNotifications(w http.ResponseWriter, r *http.Request) {
	op := "content.controller.GetNotifications"
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

	// Получение query параметров из запроса
	offsetStr := r.URL.Query().Get(offsetParam)
	limitStr := r.URL.Query().Get(limitParam)
	status := r.URL.Query().Get(statusParam)

	opt := pkgModels.NewNotificationsOpt(offsetStr, limitStr, status)

	ntfs, err := handler.serv.GetNotifications(ctx, userID, opt)
	if err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	notifications := models.MapNotifcationsServToCntrl(ntfs)

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(notifications); err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
	}
}
