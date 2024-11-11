package interafces

import (
	"context"
	bJWT "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service/jwt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
)

// AuthBehavior интерфейс с которым взаимодействует уровень controller
// Т.е. те, методы, которые нужны от service на уровне controller
type AuthBehavior interface {
	// RegisterNewUser - регистрация | добавление нового пользователя, генерация для него jwt
	RegisterNewUser(ctx context.Context, username string, password string) (bJWT.TokenString, error)

	// LoginUser - авторизация | проверяет существует ли пользователь, верный ли пароль, генерирует jwt для него
	LoginUser(ctx context.Context, username string, password string) (bJWT.TokenString, error)

	// LogoutUser - выход из системы пользователя | удаляет сессию пользователя в будущем, сейчас просто заглушка
	LogoutUser(ctx context.Context, userID models.UserID) error
}
