package controller

import (
	"net/http"

	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller/models"
	models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
)

// PostSubscriptionRealize - ручка реализации запроса пользователя на подписку на автора
func (handler *Handler) PostSubscriptionRealize(w http.ResponseWriter, r *http.Request) {
	op := "internal.account.controller.PostSubscriptionRealize"

	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Парсинг модели вводных данных из уведомления сервиса оплаты
	// ntfPayService - ответ от API сервиса оплаты
	var ntfPayService models.NotificationPaymentService
	if err := utils.ParseModels(r, &ntfPayService, op); err != nil {
		// проставляем http.StatusBadRequest
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// logger.StandardDebugF(ctx, op,
	// 	"INFO FROM YOOMONEY=%v", ntfPayService)

	// Обращение в service
	err := handler.serv.RealizeSubscriptionRequest(ctx, ntfPayService.Object.ID, ntfPayService.Object.Paid, ntfPayService.Object.Description, "")
	if err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	w.WriteHeader(http.StatusOK)
}
