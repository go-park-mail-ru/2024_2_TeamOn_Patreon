package middlewares

import "net/http"

// GetCSRFTokenHandler Handler для выдачи CSRF токена клиенту
func GetCSRFTokenHandler(w http.ResponseWriter, r *http.Request) {
	token, err := generateCSRFToken()
	if err != nil {
		http.Error(w, "Failed to generate CSRF token", http.StatusInternalServerError)
		return
	}
	// Устанавливаем CSRF токен в cookie
	http.SetCookie(w, &http.Cookie{
		Name:     csrfTokenCookieName,
		Value:    token,
		HttpOnly: true,
		Secure:   true, // Используйте true, если у вас HTTPS
	})
	// Отправляем CSRF токен в теле ответа (опционально)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"csrfToken":"` + token + `"}`))
}
