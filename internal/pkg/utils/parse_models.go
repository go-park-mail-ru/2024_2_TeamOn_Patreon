package utils

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/mailru/easyjson"
)

func ParseModels(r *http.Request, m interface{}, op string) error {
	ctx := r.Context()

	rawBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logger.StandardWarnF(ctx, op, "Failed to read request body: %v", err)
		return global.ErrInvalidJSON
	}
	defer r.Body.Close()

	// Если model реализует интерфейс easyjson.Unmarshaler
	if unmarshaler, ok := m.(easyjson.Unmarshaler); ok {
		if err := easyjson.Unmarshal(rawBytes, unmarshaler); err != nil {
			logger.StandardWarnF(ctx, op, "Received parsing with easyjson error {%v}", err)
			return global.ErrInvalidJSON
		}
		logger.StandardDebugF(ctx, op, "Parsed models with easyjson l={%v}", m)
		return nil
	}

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		logger.StandardWarnF(ctx, op, "Received parsing error {%v}", err)
		return global.ErrInvalidJSON
	}
	logger.StandardDebugF(ctx, op, "Parsed models l={%v}", m)
	return nil
}
