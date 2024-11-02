package service

import (
	bJWT "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service/jwt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service/validator"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/pkg/errors"
)

// RegisterNewUser - регистрация | добавление нового пользователя, генерация для него jwt
func (b *Behavior) RegisterNewUser(username string, password string) (bJWT.TokenString, error) {
	op := "internal.service.service.RegisterNewUsername"

	username, password, err := b.validateRegisterInput(username, password)
	if err != nil {
		return "", errors.Wrap(err, op)
	}

	// хэширование пароля
	hash, err := hashPassword(password)
	if err != nil {
		return "", errors.Wrap(err, op)
	}

	// сохранение юзера в БД и получение модельку пользователя
	user, err := b.saveUser(username, hash)
	if err != nil {
		return "", errors.Wrap(err, op)
	}

	// сгенерировать для пользователя токен
	token, err := createJWT(user)
	if err != nil {
		return "", errors.Wrap(err, op)
	}

	// вернуть токен
	return token, nil
}

func (b *Behavior) validateRegisterInput(username, password string) (string, string, error) {
	op := "internal.service.validateRegisterInput"

	username, password, err := validator.ValidateUsernameAndPassword(username, password)
	if err != nil {
		return username, password, errors.Wrap(err, op)
	}

	// Проверка есть ли такой username
	// если произошла ошибка, вернуть её
	exists, err := b.isUserExists(username)
	if err != nil {
		return username, password, errors.Wrap(err, op)
	}
	if exists {
		return username, password, errors.Wrap(global.ErrUserAlreadyExists, op)
	}
	return username, password, nil
}

func (b *Behavior) saveUser(username string, hash string) (*bModels.User, error) {
	op := "auth.service.SaveUser"

	role := bModels.Reader
	user, err := b.rep.SaveUser(username, string(role), hash)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	return user, nil
}
