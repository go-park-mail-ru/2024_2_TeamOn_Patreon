package middlewares

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service/jwt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
)

// AuthMiddleware - middleware, обрабатывает JWT токен из cookie
// передает модельку юзера в контекст
// Не возвращает ошибку если пользователь не авторизован
// просто кладет пользователя в контекст
func (m *Monster) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		op := "internal.pkg.middlewares.AuthMiddleware"

		token, err := jwt.JWTStringFromCookie(r)
		if err != nil {
			// передаем управление дальше
			next.ServeHTTP(w, r)
			return
		}

		isLogged, userID, err := m.client.VerifyToken(token)

		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		if !isLogged {
			next.ServeHTTP(w, r)
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
