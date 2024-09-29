package jwt

import (
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/buisness/models"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// ПОКА ТАК, потом будем парсить из конфига
var jwtKey = []byte("secret-key-456764459876")

// CreateJWT
// Функция создания JWT токена по данным
// ttl - время жизни в ЧАСАХ
func CreateJWT(user bModels.User, ttl int) (TokenString, error) {
	// по умолчанию 24 часа
	if ttl == 0 {
		ttl = 24
	}

	// создаем структуру токена claims
	claims := TokenClaims{
		UserID:   user.UserID,
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(ttl))), // Токен истекает через ttl часа
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                     // Время выпуска токена
			Issuer:    "auth-app",                                                         // Кто выпустил токен
		},
	}

	// Создаем токен
	tokenString, err := createToken(claims)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// createToken
// Создает токен по структуре TokenClaims
// также подписывает токен
func createToken(tokenClaims TokenClaims) (TokenString, error) {
	// Создаем токен с помощью NewWithClaims, передавая наш объект tokenClaims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	// Подписываем токен с использованием секретного ключа
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return TokenString(tokenString), nil
}
