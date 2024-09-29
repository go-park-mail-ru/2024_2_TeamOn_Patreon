package utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/errors"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/logger"
	"net/http"
)

func ParseModels(r *http.Request, m any, op string) *errors.MsgError {
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		logger.StandardWarn(fmt.Sprintf("Resived parsing error {%v}", err), op)
		return errors.UnknownError(err, op)
	}
	logger.StandardDebug(fmt.Sprintf("Parsed models l={%v}", m), op)
	return nil
}
