/*
** HOW TO USE ??

Предлагаю использовать при создании роутера

Это обертка над HandlerFunc, что бы запросы логгировались
возвращает HandlerFunc!!

Пример использования:
```
newHandlerFunc := Logging(handlerFuncWithOutLogger)
```

! Не забудьте проимпортировать
*/

package utils

import (
	"fmt"
	"log/slog"
	"net/http"
)

func Logging(handler http.Handler, op string) http.Handler {
	newHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info(fmt.Sprintf("new message %s %s | in %s", r.Method, r.URL.Path, op))
		handler.ServeHTTP(w, r)
	})
	return newHandler
}
