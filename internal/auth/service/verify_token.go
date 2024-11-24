package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service/jwt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service/mapper"
	"github.com/pkg/errors"
)

// VerifyToken - проверяет токен ан валидность и на просрок
// Возвращает инфу валидный ли токен и uuid пользователя
func (b *Behavior) VerifyToken(ctx context.Context, token string) (isLogged bool, userID string, err error) {
	op := "auth.service.VerifyToken"

	// TODO: проверка не протух ли токен

	claims, err := jwt.ParseJWTFromJWTString(token)
	if err != nil {
		return false, "", errors.Wrap(err, op)
	}

	user := mapper.MapTokenToUser(claims)
	return true, string(user.UserID), nil
}
