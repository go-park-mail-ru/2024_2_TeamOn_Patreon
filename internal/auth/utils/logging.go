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
