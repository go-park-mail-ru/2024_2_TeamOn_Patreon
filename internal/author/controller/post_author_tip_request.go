package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller/models"
	cModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/controller/models"
	models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/controller/models"
	global "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/gorilla/mux"

	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
)

// PostAuthorTip - ручка создания запроса на пожертвование автору
func (handler *Handler) PostAuthorTip(w http.ResponseWriter, r *http.Request) {
	op := "internal.account.controller.PostAuthorTip"

	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Определяем authorID
	vars := mux.Vars(r)
	authorID := vars["authorID"]

	// Извлекаем userID из контекста
	userData, ok := r.Context().Value(global.UserKey).(bModels.User)
	if !ok {
		logger.StandardResponse(ctx, "userData not found in context", http.StatusUnauthorized, r.Host, op)
		// Status 401
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	userID := string(userData.UserID)

	// Если пользователь пытается задонатить себе
	if userID == authorID {
		logger.StandardResponse(ctx, "user can't donate to himself", http.StatusBadRequest, r.Host, op)
		// Status 400
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Валидация userID на соответствие стандарту UUIDv4
	if ok := utils.IsValidUUIDv4(userID); !ok {
		// Status 400
		logger.StandardResponse(ctx, "invalid userID format", http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Валидация authorID на соответствие стандарту UUIDv4
	if ok := utils.IsValidUUIDv4(authorID); !ok {
		// Status 400
		logger.StandardResponse(ctx, "invalid authorID format", http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Парсинг данных из json
	tipInfo := &cModels.Tip{}
	if err := json.NewDecoder(r.Body).Decode(&tipInfo); err != nil {
		logger.StandardWarnF(ctx, op, "json parsing error {%v}", err)
		// Status 400
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Валидация суммы (не меньше 10 р)
	if tipInfo.Cost < 10 {
		logger.StandardWarnF(ctx, op, "the donation amount cannot be less than 10 rubles")
		// Status 400
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Собираем информацию для пользователя о платеже
	strCost := strconv.Itoa(tipInfo.Cost) + ".00"
	payInfo := models.InfoPaySubscription{
		AuthorID:    authorID,
		Cost:        strCost,
		Description: fmt.Sprintf("Пожертвование на сумму %v руб.", strCost),
		PayType:     models.TypeTip,
	}

	// Обращение к API оплаты
	paymentResponse, err := handler.CreateRequestPay(ctx, payInfo)
	if err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Обращение к service
	tipReq := models.TipRequest{
		TipReqID: paymentResponse.ID,
		UserID:   userID,
		AuthorID: authorID,
		Cost:     tipInfo.Cost,
		Message:  tipInfo.Message,
	}

	if err := handler.serv.CreateTipRequest(r.Context(), models.MapControllerTipReqToServTipReq(tipReq)); err != nil {
		logger.StandardWarnF(ctx, op, "update info error {%v}", err)
		// Status 500
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Возвращаем URL на API оплаты
	if err = json.NewEncoder(w).Encode(paymentResponse.Confirmation.ConfirmationURL); err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
	}

	// Status 200
	w.WriteHeader(http.StatusOK)
}
