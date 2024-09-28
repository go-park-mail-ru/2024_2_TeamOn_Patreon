package jwt

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/global"
	"github.com/golang-jwt/jwt/v5"
)

// ParseJWTFromCookie
// Функция для парсинга и верификации JWT токена из Cookie
// Также валидирует данные
func ParseJWTFromCookie(r *http.Request) (*TokenClaims, error) {
	// Ищем куку с токеном
	cookie, err := r.Cookie(global.CookieJWT)
	if err != nil {
		return nil, fmt.Errorf("cookie 'token' not found: %v", err)
	}

	// Извлекаем JWT токен из значения куки
	tokenString := cookie.Value

	claims, err := ParseJWTFromJWTString(tokenString)

	if err != nil {
		return nil, err
	}

	// Токен успешно проверен, возвращаем распарсенные данные
	return claims, nil
}

func ParseJWTFromJWTString(tokenString string) (*TokenClaims, error) {
	// Объект для хранения наших данных из токена
	claims := &TokenClaims{}

	// Парсим токен, проверяя его подпись и поля
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что метод подписи — HMAC, и возвращаем секретный ключ
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	// Проверяем, валиден ли токен
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	// Проверяем валидность токена и его полей
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Проверка на истечение времени действия токена
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, fmt.Errorf("token has expired")
	}

	// Токен успешно проверен, возвращаем распарсенные данные
	return claims, nil
}
