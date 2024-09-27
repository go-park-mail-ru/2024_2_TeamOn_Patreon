package errors

import (
	"fmt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/polina-auth/internal/common/logger"
)

func UnknownError(err error, op string) *MsgError {
	errMsg := "неизвестная ошибка"
	if err == nil {
		err = fmt.Errorf("unknown error in %v", op)
	}
	logger.StandardError(err.Error(), op)
	return New(err.Error(), errMsg)
}
