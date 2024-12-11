package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

const (
	apiURL    = "https://api.yookassa.ru/v3/payments"              // URL для создания платежа
	clientID  = "996895"                                           // СПРЯТАТЬ!!!
	secretKey = "test_kKJUwERP7sXkFVyy1mrjp-82dg8-bohbwnsodUk3peA" // СПРЯТАТЬ!!!
)

type PaymentRequest struct {
	Amount       Amount       `json:"amount"`
	Description  string       `json:"description"`
	Confirmation Confirmation `json:"confirmation"`
	Metadata     interface{}  `json:"metadata,omitempty"`
	Test         bool         `json:"test"`
}

type Amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type Confirmation struct {
	Type      string `json:"type"`
	ReturnURL string `json:"return_url,omitempty"`
}

type PaymentResponse struct {
	ID          string `json:"id"`
	Status      string `json:"status"`
	Paid        bool   `json:"paid"`
	Amount      Amount `json:"amount"`
	CreatedAt   string `json:"created_at"`
	Description string `json:"description"`
}

func main() {
	paymentRequest := PaymentRequest{
		Amount: Amount{
			Value:    "199999.00",
			Currency: "RUB",
		},
		Description: "Лиле на мильён вкусняшек",
		Confirmation: Confirmation{
			Type:      "redirect",
			ReturnURL: "https://pushart.online/profile/da2560c6-9c85-4b40-916d-dd42db575112", // Куда направить после оплаты
		},
		Test: true,
	}

	// Преобразуем запрос в JSON
	requestBody, err := json.Marshal(paymentRequest)
	if err != nil {
		fmt.Println("Ошибка при преобразовании в JSON:", err)
		return
	}

	// Создаем HTTP-запрос
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return
	}

	// Идемпотент
	idempotenceKey := uuid.New().String()

	// Устанавливаем заголовки
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+basicAuth(clientID, secretKey))
	req.Header.Set("Idempotence-Key", idempotenceKey) // Генерируем уникальный ключ

	// Отправляем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	// Выводим статус и тело ответа
	fmt.Println("Статус ответа:", resp.Status)
	fmt.Println("Тело ответа:", string(body))

	// Проверяем статус платежа
	var paymentResponse PaymentResponse
	if err := json.Unmarshal(body, &paymentResponse); err != nil {
		fmt.Println("Ошибка при разборе ответа:", err)
		return
	}

	payID := "2eebdfa1-000f-5000-a000-1de5ba31f906"
	// Вызов функции для проверки статуса платежа
	checkPaymentStatus(payID)
}

// Функция для базовой аутентификации
func basicAuth(clientID, secretKey string) string {
	auth := clientID + ":" + secretKey
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// Функция для проверки статуса платежа
func checkPaymentStatus(paymentID string) {
	apiURL := fmt.Sprintf("https://api.yookassa.ru/v3/payments/%s", paymentID)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return
	}

	req.Header.Set("Authorization", "Basic "+basicAuth(clientID, secretKey))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	// Выводим статус и тело ответа
	fmt.Println("Статус ответа проверки:", resp.Status)
	fmt.Println("Тело ответа проверки:", string(body))

	// Обработка ответа
	var paymentStatusResponse PaymentResponse
	if err := json.Unmarshal(body, &paymentStatusResponse); err != nil {
		fmt.Println("Ошибка при разборе ответа:", err)
		return
	}

	// Выводим информацию о статусе платежа
	fmt.Println("ID платежа:", paymentStatusResponse.ID)
	fmt.Println("Статус платежа:", paymentStatusResponse.Status)
	fmt.Println("Оплачен:", paymentStatusResponse.Paid)
	fmt.Println("Сумма:", paymentStatusResponse.Amount.Value, paymentStatusResponse.Amount.Currency)
	fmt.Println("Описание:", paymentStatusResponse.Description)
	fmt.Println("Дата создания:", paymentStatusResponse.CreatedAt)
}
