package behavior

import (
	hasher "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/behavior/hasher"
	bJWT "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/behavior/jwt"
	session "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/behavior/session"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/config"
	rInterfaces "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/repository/interfaces"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/business/models"
	global "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/global"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/logger"
	"github.com/pkg/errors"
)

type Behavior struct {
	rep rInterfaces.AuthRepository
}

func New(repository rInterfaces.AuthRepository) *Behavior {
	return &Behavior{repository}
}

// RegisterNewUser - регистрация | добавление нового пользователя, генерация для него jwt
func (b *Behavior) RegisterNewUser(username string, password string) (string, error) {
	op := "internal.behavior.behavior.RegisterNewUsername"

	// Проверка есть ли такой username
	// если произошла ошибка, вернуть её
	exists, err := b.isUserExists(username)
	if err != nil {
		return "", errors.Wrap(err, op)
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

	// создаем сессию
	sessionModel := session.CreateSession(user.UserID)

	// сохраняем сессию
	sessionString, err := b.saveSession(sessionModel)
	if err != nil {
		return "", errors.Wrap(err, op)
	}

	// возвращаем sessionID
	return sessionString, nil
}

// AuthoriseUser - авторизация | проверяет существует ли пользователь, верный ли пароль, генерирует jwt для него
func (b *Behavior) AuthoriseUser(username string, password string) (string, error) {
	op := "auth.api.AuthoriseUser"

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

	//// сгенерировать для пользователя токен
	//token, err := createJWT(user)
	//if err != nil {
	//	logger.StandardDebugF(op, "Authorisation failed: user %s generation token failed", username)
	//	return "", errors.Wrap(err, op)
	//}

	// создаем сессию
	sessionModel := session.CreateSession(user.UserID)

	// сохраняем сессию
	sessionString, err := b.saveSession(sessionModel)
	if err != nil {
		return "", errors.Wrap(err, op)
	}

	logger.StandardDebugF(op, "Login user={%v} with token={%v}", username, sessionString)

	return sessionString, nil
}

func (b *Behavior) isUserExists(username string) (bool, error) {
	op := "auth.behavior.IsUserExists"

	// Проверка есть ли такой username
	// если произошла ошибка, вернуть её
	exists, err := b.rep.UserExists(username)
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	return exists, nil
}

func hashPassword(password string) (string, error) {
	op := "auth.behavior.HashPassword"

	hash, err := hasher.HashPassword(password)
	if err != nil {
		return "", errors.Wrap(config.ErrServer, op)
	}
	return hash, nil
}

func (b *Behavior) saveUser(username string, hash string) (*bModels.User, error) {
	op := "auth.behavior.SaveUser"

	role := bModels.Reader
	user, err := b.rep.SaveUser(username, int(role), hash)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	return user, nil
}

func createJWT(user *bModels.User) (bJWT.TokenString, error) {
	op := "auth.behavior.CreateJWT"

	// сгенерировать для пользователя токен
	token, err := bJWT.CreateJWT(*user, global.TTL)
	if err != nil {
		return "", errors.Wrap(err, op)
	}
	return token, nil
}

func (b *Behavior) getUser(username string) (*bModels.User, error) {
	op := "auth.behavior.GetUser"

	user, err := b.rep.GetUserByUsername(username)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	return user, nil
}

func (b *Behavior) comparePassword(user bModels.User, password string) (bool, error) {
	op := "auth.behavior.ComparePassword"

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

//func (b *Behavior) DeleteUser(username string) error {
//	op := "auth.behavior.DeleteUser"
//	user, err := b.rep.GetUserByUsername(username)
//	if err != nil {
//		return errors.Wrap(err, op)
//	}
//
//}

func (b *Behavior) saveSession(sessionModel *session.SessionModel) (string, error) {
	op := "auth.behavior.SaveSession"

	sessionString, err := b.rep.SaveSession(*sessionModel)
	if err != nil {
		return "", errors.Wrap(err, op)
	}
	return sessionString, nil
}
