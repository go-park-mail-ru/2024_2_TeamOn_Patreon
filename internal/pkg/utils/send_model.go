package utils

import (
	"context"
	"encoding/json"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"net/http"
)

func SendModel(model any, w http.ResponseWriter, op string, ctx context.Context) {
	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Сериализуем модель в JSON и отправляем в ответ
	json.NewEncoder(w).Encode(model)
	logger.StandardDebugF(ctx, op, "Sent model = %v", model)
}
