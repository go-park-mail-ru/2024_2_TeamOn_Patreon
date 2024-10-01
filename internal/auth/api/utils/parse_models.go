package utils

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/config"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/logger"
	"net/http"
)

func ParseModels(r *http.Request, m any, op string) error {
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		logger.StandardWarnF(op, "Resived parsing error {%v}", err)
		return config.ErrInvalidJSON
	}
	logger.StandardDebugF(op, "Parsed models l={%v}", m)
	return nil
}
