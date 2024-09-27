package utils

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/polina-auth/internal/auth/behavior/jwt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/polina-auth/internal/common/global"
	"net/http"
	"time"
)

func CreateCookie(tokenString jwt.TokenString) http.Cookie {
	// Сохранение токена в куки
	// Устанавливаем токен в куку
	expirationTime := time.Now().Add(global.TTL * time.Hour) // Время жизни куки должно совпадать со временем жизни токена
	cookie := http.Cookie{
		Name:     global.CookieJWT,    // Имя куки
		Value:    string(tokenString), // Значение куки — наш сгенерированный токен
		Expires:  expirationTime,      // Время истечения куки
		HttpOnly: true,                // Кука доступна только через HTTP (не через JS)
		Path:     "/",
		// Domain:   "example.com",       // Домен, на котором будет действовать кука
		// Secure:   true,                // Кука передаётся только по HTTPS
	}
	return cookie
}
