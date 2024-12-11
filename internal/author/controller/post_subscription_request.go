package controller

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller/models"
	cModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/controller/models"

	models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/pkg/errors"
)

const (
	apiURL    = "https://api.yookassa.ru/v3/payments"              // URL для создания платежа
	clientID  = "996895"                                           // СПРЯТАТЬ!!!
	secretKey = "test_kKJUwERP7sXkFVyy1mrjp-82dg8-bohbwnsodUk3peA" // СПРЯТАТЬ!!!
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

	// Получение через service суммы оплаты
	costSubscription, err := handler.serv.GetCostSubscription(ctx, subReq.MonthCount, subReq.AuthorID, subReq.Layer)
	if err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	payInfo := models.InfoPaySubscription{
		AuthorID:    subReq.AuthorID,
		Cost:        costSubscription,
		Description: fmt.Sprintf("Оформление подписки на %v мес.", subReq.MonthCount),
	}

	// Обращение к API оплаты
	paymentResponse, err := handler.createRequestPay(ctx, payInfo)
	if err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Обращение в service
	subReq.SubscriptionRequestID = paymentResponse.ID
	err = handler.serv.CreateSubscriptionRequest(ctx, models.MapSubReqToServiceSubReq(userID, subReq))
	if err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Возвращаем URL на API оплаты
	if err = json.NewEncoder(w).Encode(paymentResponse.Confirmation.ConfirmationURL); err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *Handler) createRequestPay(ctx context.Context, payInfo models.InfoPaySubscription) (cModels.PaymentResponse, error) {
	op := "internal.author.controller.createRequestPay"

	// Модель запроса к сервису оплаты с указанием суммы и редиректа после успешной оплаты
	paymentRequest := cModels.PaymentRequest{
		Amount: cModels.Amount{
			Value:    payInfo.Cost,
			Currency: "RUB",
		},
		Description: payInfo.Description,
		Confirmation: cModels.ConfirmationReq{
			Type:      "redirect",
			ReturnURL: "https://pushart.online/profile/" + payInfo.AuthorID, // Куда направить после оплаты
		},
		Test: true,
	}

	var paymentResponse cModels.PaymentResponse

	// Преобразуем запрос в JSON
	requestBody, err := json.Marshal(paymentRequest)
	if err != nil {
		logger.StandardDebugF(ctx, op, "marshal err: %v", err)
		return paymentResponse, errors.Wrap(err, op)
	}

	// Создаем HTTP-запрос
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		logger.StandardDebugF(ctx, op, "Fail create request: %v", err)
		return paymentResponse, errors.Wrap(err, op)
	}

	// Идемпотент
	idempotenceKey := utils.GenerateUUID()

	// Устанавливаем заголовки
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+basicAuth(clientID, secretKey))
	req.Header.Set("Idempotence-Key", idempotenceKey) // Генерируем уникальный ключ

	// Отправляем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.StandardDebugF(ctx, op, "Fail send request: %v", err)
		return paymentResponse, errors.Wrap(err, op)
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.StandardDebugF(ctx, op, "Fail read response: %v", err)
		return paymentResponse, errors.Wrap(err, op)
	}

	// Выводим статус и тело ответа
	logger.StandardInfo(ctx,
		fmt.Sprintf("Статус ответа: %v. Тело ответа: %v", resp.Status, string(body)), op)

	// Получаем ID платежа и URL на оплату
	if err := json.Unmarshal(body, &paymentResponse); err != nil {
		logger.StandardDebugF(ctx, op, "Fail unmarshal response: %v", err)
		return paymentResponse, errors.Wrap(err, op)
	}

	return paymentResponse, nil
}

func basicAuth(clientID, secretKey string) string {
	auth := clientID + ":" + secretKey
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
