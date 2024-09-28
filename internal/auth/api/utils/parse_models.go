package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/errors"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/logger"
)

func ParseModels(r *http.Request, m any, op string) *errors.MsgError {
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		return errors.UnknownError(err, op)
	}
	logger.StandardDebug(
		fmt.Sprintf("ParseModels (%v)", m),
		op)
	return nil
}
