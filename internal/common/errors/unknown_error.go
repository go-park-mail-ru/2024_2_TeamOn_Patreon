package errors

import (
	"fmt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/logger"
	"net/http"
)

func UnknownError(err error, op string) *MsgError {
	errMsg := "неизвестная ошибка"
	if err == nil {
		err = fmt.Errorf("unknown error in %v", op)
	}
	logger.StandardError(err.Error(), op)
	return NewCode(err.Error(), errMsg, http.StatusInternalServerError)
}
