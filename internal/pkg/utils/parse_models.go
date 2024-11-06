package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
)

func ParseModels(r *http.Request, m any, op string) error {
	ctx := r.Context()
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		logger.StandardWarnF(ctx, op, "Resived parsing error {%v}", err)
		return global.ErrInvalidJSON
	}
	logger.StandardDebugF(ctx, op, "Parsed models l={%v}", m)
	return nil
}
