package utils

import (
	"fmt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/behavior"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/errors"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/logger"
	"net/http"
)

func GetBehaviorCtx(r *http.Request, op string) (*behavior.Behavior, *errors.MsgError) {
	// Получаем Behavior из контекста
	bInterface := r.Context().Value(global.BehaviorKey)
	if bInterface == nil {
		errM := errors.UnknownError(fmt.Errorf("Received empty behavior"), op)
		logger.StandardError(errM.Error(), op)
		return nil, errM
	}

	b, ok := bInterface.(*behavior.Behavior)
	if !ok {
		errM := errors.UnknownError(fmt.Errorf("Received invalid behavior"), op)
		logger.StandardError(errM.Error(), op)
		return nil, errM
	}
	logger.StandardDebugF(op, "Getted behavior={%v}", b)
	return b, nil
}
