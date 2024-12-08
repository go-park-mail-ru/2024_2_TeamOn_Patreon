package middlewares

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"net/http"

	"github.com/pborman/uuid"
)

func AddRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := uuid.NewRandom().String()                            // Генерация уникального ID
		ctx := context.WithValue(r.Context(), global.CtxReqId, reqID) // Добавление в контекст
		next.ServeHTTP(w, r.WithContext(ctx))                         // Передача контекста дальше
	})
}
