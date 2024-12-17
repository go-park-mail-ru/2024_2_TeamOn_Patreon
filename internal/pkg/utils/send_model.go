package utils

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/mailru/easyjson"
)

func SendModel(model easyjson.Marshaler, w http.ResponseWriter, op string, ctx context.Context) {
	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Используем easyjson для сериализации
	rawBytes, err := easyjson.Marshal(model)
	if err != nil {
		logger.StandardDebugF(ctx, op, "Failed to marshal model with easyjson: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Записываем сериализованные данные в ResponseWriter
	if _, err := w.Write(rawBytes); err != nil {
		logger.StandardDebugF(ctx, op, "Failed to write response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.StandardDebugF(ctx, op, "Sent model = %v", model)
}
