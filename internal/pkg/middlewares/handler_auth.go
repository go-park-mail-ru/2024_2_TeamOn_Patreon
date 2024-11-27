package middlewares

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service/jwt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
)

// HandlerAuth - middleware, обрабатывает JWT токен из cookie
// передает модельку юзера в контекст
// Возвращает ошибку, если пользователь не авторизован!
func (m *Monster) HandlerAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		op := "internal.pkg.middlewares.HandlerAuth"

		token, err := jwt.JWTStringFromCookie(r)
		if err != nil {
			logger.StandardDebugF(r.Context(), op, "Auth failed: fail get token err=%v", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		isLogged, userID, err := m.client.VerifyToken(token)

		if err != nil {
			logger.StandardDebugF(r.Context(), op, "Auth failed: fail get user from token err=%v with grpc", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if !isLogged {
			logger.StandardDebugF(r.Context(), op, "Auth failed: user is not logged")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		user := models.User{
			UserID: models.UserID(userID),
		}

		// передаем в контекст
		ctx := context.WithValue(r.Context(), global.UserKey, user)
		logger.StandardDebugF(r.Context(), op, "Transferred user (id={%v}) in ctx",
			user.UserID)

		// добавляем контекст в контекст r
		r = r.WithContext(ctx)

		// передаем управление дальше
		next.ServeHTTP(w, r)
	})
}
