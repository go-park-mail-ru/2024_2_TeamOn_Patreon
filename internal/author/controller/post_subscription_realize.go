package controller

import (
	"net/http"

	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller/models"
	models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
)

// PostSubscriptionRealize - ручка реализации запроса пользователя на подписку на автора
func (handler *Handler) PostSubscriptionRealize(w http.ResponseWriter, r *http.Request) {
	op := "internal.account.controller.PostSubscriptionRealize"

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

	subReqID := subReq.SubscriptionRequestID

	// Валидация введённых данных
	if ok := utils.IsValidUUIDv4(subReqID); !ok {
		err := global.ErrIsInvalidUUID
		logger.StandardWarnF(ctx, op, "subReqID={%v} is not UUID", userID)
		// Status 400
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Обращение в service
	err := handler.serv.RealizeSubscriptionRequest(ctx, subReqID)
	if err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	w.WriteHeader(http.StatusOK)
}
