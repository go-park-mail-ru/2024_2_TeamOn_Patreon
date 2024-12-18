/*
** HOW TO USE ??

Предлагаю использовать при создании роутера

Это обертка над HandlerFunc, что бы запросы логгировались
возвращает HandlerFunc!!

Пример использования:
```
router.Use(middlewares.Logging)
```

! Не забудьте проимпортировать
*/

package middlewares

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/gorilla/mux"
)

func Logging(handler http.Handler) http.Handler {
	op := "pkg.middlewares.Logging"

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		if route != nil {
			op = route.GetName()
		}
		logger.StandardInfoF(r.Context(), op, "Received request %s %s", r.Method, r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}
