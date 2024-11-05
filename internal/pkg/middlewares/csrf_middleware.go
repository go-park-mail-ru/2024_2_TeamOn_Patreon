package middlewares

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
)

const csrfTokenCookieName = "csrf_token_yd23g8763dfg28cg2e3iuy"

// generateCSRFToken создает случайный токен для защиты от CSRF
func generateCSRFToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// CsrfMiddleware для проверки CSRF токена
func CsrfMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodDelete {
			// Извлекаем CSRF токен из заголовка запроса
			csrfToken := r.Header.Get("X-CSRF-Token")
			// Получаем токен из cookie
			cookie, err := r.Cookie(csrfTokenCookieName)
			if err != nil || csrfToken != cookie.Value {
				http.Error(w, "CSRF token mismatch", http.StatusForbidden)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
