package utils

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/errors"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/logger"
	"net/http"
)

func ParseModels(r *http.Request, m any, op string) *errors.MsgError {
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		logger.StandardWarnF(op, "Resived parsing error {%v}", err)
		return errors.UnknownError(err, op)
	}
	logger.StandardDebugF(op, "Parsed models l={%v}", m)
	return nil
}
