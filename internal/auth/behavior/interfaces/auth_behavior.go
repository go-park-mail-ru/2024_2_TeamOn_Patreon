package interfaces

import (
	bJWT "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/behavior/jwt"
	cErrors "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/errors"
)

type AuthBehavior interface {
	// RegisterNewUser - регистрация | добавление нового пользователя, генерация для него jwt
	RegisterNewUser(username string, password string) (bJWT.TokenString, *cErrors.MsgError)

	// AuthoriseUser - авторизация | проверяет существует ли пользователь, верный ли пароль, генерирует jwt для него
	AuthoriseUser(username string, password string) (bJWT.TokenString, *cErrors.MsgError)
}