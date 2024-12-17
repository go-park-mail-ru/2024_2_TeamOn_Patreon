package utils

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/mailru/easyjson"
)

func SendModel(model interface{}, w http.ResponseWriter, op string, ctx context.Context) {
	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Если model реализует интерфейс easyjson.Marshaler
	if marshaler, ok := model.(easyjson.Marshaler); ok {
		// Используем easyjson для сериализации
		rawBytes, err := easyjson.Marshal(marshaler)
		if err != nil {
			logger.StandardDebugF(ctx, op, "Failed to marshal model with easyjson: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(rawBytes); err != nil {
			logger.StandardDebugF(ctx, op, "Failed to write response: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		logger.StandardDebugF(ctx, op, "Sent model with easyjson = %v", model)
		return
	}

	// Если не реализует easyjson.Marshaler, используем стандартный Encoder
	json.NewEncoder(w).Encode(model)
	logger.StandardDebugF(ctx, op, "Sent model = %v", model)
}
