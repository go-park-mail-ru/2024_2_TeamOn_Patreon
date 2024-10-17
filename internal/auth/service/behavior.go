package service

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/config"
	hasher "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service/hasher"
	rInterfaces "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service/interfaces"
	bJWT "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service/jwt"
	global "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/pkg/errors"
)

type Behavior struct {
	rep rInterfaces.AuthRepository
}

func New(repository rInterfaces.AuthRepository) *Behavior {
	return &Behavior{repository}
}

// RegisterNewUser - регистрация | добавление нового пользователя, генерация для него jwt
func (b *Behavior) RegisterNewUser(username string, password string) (bJWT.TokenString, error) {
	op := "internal.service.service.RegisterNewUsername"

	// Проверка есть ли такой username
	// если произошла ошибка, вернуть её
	exists, errM := b.isUserExists(username)
	if errM != nil {
		return "", errors.Wrap(errM, op)
	}
	if exists {
		return "", errors.Wrap(config.ErrUserAlreadyExists, op)
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

// AuthoriseUser - авторизация | проверяет существует ли пользователь, верный ли пароль, генерирует jwt для него
func (b *Behavior) AuthoriseUser(username string, password string) (bJWT.TokenString, error) {
	op := "auth.controller.AuthoriseUser"

	// проверяем существует ли пользователь
	exists, err := b.isUserExists(username)
	// если не существует или какая-то ошибка, выходим
	if err != nil {
		logger.StandardDebugF(op, "Authorisation failed: user %s does not exist or err", username)
		return "", errors.Wrap(err, op)
	}
	if !exists {
		logger.StandardDebugF(op, "Authorisation failed: user %s does not exist", username)
		return "", config.ErrNotValidUserAndPassword
	}

	// получаем модельку юзера по username
	user, err := b.getUser(username)
	if err != nil {
		logger.StandardDebugF(op, "Authorisation failed: user %s does not exist", username)
		return "", errors.Wrap(err, op)
	}

	// сравниваем пароли
	ok, err := b.comparePassword(*user, password)
	if err != nil {
		logger.StandardDebugF(op, "Authorisation failed: user %s does not match", username)
		return "", errors.Wrap(err, op)
	}
	if !ok {
		logger.StandardDebugF(op, "Authorisation failed: user %s does not match", username)
		return "", errors.Wrap(config.ErrNotValidUserAndPassword, op)
	}

	// сгенерировать для пользователя токен
	token, err := createJWT(user)
	if err != nil {
		logger.StandardDebugF(op, "Authorisation failed: user %s generation token failed", username)
		return "", errors.Wrap(err, op)
	}

	logger.StandardDebugF(op, "Login user={%v} with token={%v}", username, token)

	return token, nil
}

func (b *Behavior) isUserExists(username string) (bool, error) {
	op := "auth.service.IsUserExists"

	// Проверка есть ли такой username
	// если произошла ошибка, вернуть её
	exists, err := b.rep.UserExists(username)
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	return exists, nil
}

func hashPassword(password string) (string, error) {
	op := "auth.service.HashPassword"

	hash, err := hasher.HashPassword(password)
	if err != nil {
		return "", errors.Wrap(config.ErrServer, op)
	}
	return hash, nil
}

func (b *Behavior) saveUser(username string, hash string) (*bModels.User, error) {
	op := "auth.service.SaveUser"

	role := bModels.Reader
	user, err := b.rep.SaveUser(username, int(role), hash)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	return user, nil
}

func createJWT(user *bModels.User) (bJWT.TokenString, error) {
	op := "auth.service.CreateJWT"

	// сгенерировать для пользователя токен
	token, err := bJWT.CreateJWT(*user, global.TTL)
	if err != nil {
		return "", errors.Wrap(err, op)
	}
	return token, nil
}

func (b *Behavior) getUser(username string) (*bModels.User, error) {
	op := "auth.service.GetUser"

	user, err := b.rep.GetUserByUsername(username)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	return user, nil
}

func (b *Behavior) comparePassword(user bModels.User, password string) (bool, error) {
	op := "auth.service.ComparePassword"

	userHash, err := b.rep.GetPasswordHashByID(user.UserID)
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	// Сравниваем введённый пароль с сохранённым хэшем
	err = hasher.CheckPasswordHash(password, userHash)
	if err != nil {
		return false, errors.Wrap(config.ErrNotValidUserAndPassword, op)
	} else {
		return true, nil
	}
}

func (b *Behavior) LogoutUser(userID bModels.UserID) error {
	op := "auth.service.behavior.LogoutUser"

	logger.StandardInfoF(op, "LogoutUser user={%v}", userID)
	return nil
}
