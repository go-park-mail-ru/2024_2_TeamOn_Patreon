package controller

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	cModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/controller/models"
	models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/controller/models"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/pkg/errors"
)

const (
	apiURL    = "https://api.yookassa.ru/v3/payments"              // URL для создания платежа
	clientID  = "996895"                                           // СПРЯТАТЬ!!!
	secretKey = "test_kKJUwERP7sXkFVyy1mrjp-82dg8-bohbwnsodUk3peA" // СПРЯТАТЬ!!!
)

func (handler *Handler) CreateRequestPay(ctx context.Context, payInfo models.InfoPaySubscription) (cModels.PaymentResponse, error) {
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
