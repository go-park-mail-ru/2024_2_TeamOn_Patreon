package controller

import (
	"encoding/json"
	"net/http"

	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller/models"
	models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
)

// PostSubscriptionRequest - ручка создания запроса пользователя на подписку на автора
func (handler *Handler) PostSubscriptionRequest(w http.ResponseWriter, r *http.Request) {
	op := "internal.account.controller.PostSubscription"

	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Извлекаем userID из контекста
	userData, ok := r.Context().Value(global.UserKey).(bModels.User)
	if !ok {
		logger.StandardResponse(ctx, "userData not found in context", http.StatusUnauthorized, r.Host, op)
		// Status 401
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	userID := string(userData.UserID)

	// Валидация userID на соответствие стандарту UUIDv4
	if ok := utils.IsValidUUIDv4(userID); !ok {
		err := global.ErrIsInvalidUUID
		logger.StandardWarnF(ctx, op, "userID={%v} is not UUID", userID)
		// Status 400
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Парсинг модели вводных данных для сохранения запроса на подписку
	var subReq models.SubscriptionRequest
	if err := utils.ParseModels(r, &subReq, op); err != nil {
		// проставляем http.StatusBadRequest
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Валидация введённых данных о подписке
	if _, err := subReq.Validate(); err != nil {
		logger.StandardWarnF(ctx, op, "Received validation error={%v}", err)
		// проставляем http.StatusBadRequest
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Обращение в service
	subReqID, err := handler.serv.CreateSubscriptionRequest(ctx, models.MapSubReqToServiceSubReq(userID, subReq))
	if err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Возвращаем ID запроса на подписку
	subReq.SubscriptionRequestID = subReqID
	if err = json.NewEncoder(w).Encode(subReq.SubscriptionRequestID); err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
	}

	w.WriteHeader(http.StatusOK)
}
