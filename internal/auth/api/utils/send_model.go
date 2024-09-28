package utils

import (
	"encoding/json"
	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/polina-auth/internal/auth/api/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/polina-auth/internal/common/logger"
	"net/http"
)

func SendModel(model tModels.Model, w http.ResponseWriter, op string) {
	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Сериализуем модель в JSON и отправляем в ответ
	json.NewEncoder(w).Encode(model)
	logger.StandardSendModel(model.String(), op)
}
