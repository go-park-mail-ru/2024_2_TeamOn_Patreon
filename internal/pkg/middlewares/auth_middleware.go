package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service/jwt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service/mapper"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
)

// AuthMiddleware - middleware, обрабатывает JWT токен из cookie
// передает модельку юзера в контекст
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		op := "internal.pkg.middlewares.AuthMiddleware"

		// парсинг jwt токена
		tokenClaims, err := jwt.ParseJWTFromCookie(r)
		//if err != nil || tokenClaims == nil {
		//	w.WriteHeader(http.StatusUnauthorized)
		//	return
		//}

		if err == nil {

			// если все ок достаем юзер ид, юзернэйм и роль
			// мапим это все в структуру user для бизнес-логики
			user := mapper.MapTokenToUser(tokenClaims)

			// передаем в контекст
			ctx := context.WithValue(r.Context(), global.UserKey, user)
			logger.StandardDebug(
				r.Context(),
				fmt.Sprintf("Transferred user (id={%v}, name={%v}) in ctx", user.UserID, user.Username),
				op,
			)

			// добавляем контекст в контекст r
			r = r.WithContext(ctx)
		}

		// передаем управление дальше
		next.ServeHTTP(w, r)
	})
}
