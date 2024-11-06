package utils

import (
	"encoding/json"
	"net/http"

	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller/models"
)

func SendStringModel(model tModels.Model, w http.ResponseWriter, op string) {
	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Сериализуем модель в JSON и отправляем в ответ
	json.NewEncoder(w).Encode(model)
	// logger.StandardSendModel(model.String(), op)
}

func SendModel(model any, w http.ResponseWriter, op string) {
	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Сериализуем модель в JSON и отправляем в ответ
	json.NewEncoder(w).Encode(model)
}
