package behavior

import (
	"fmt"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/behavior/hasher"
	bJWT "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/behavior/jwt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/repository/interfaces"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/buisness/models"
	cErrors "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/errors"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/global"
	absInterface "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/interfaces"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/logger"
)

type Behavior struct {
	rep interfaces.AuthRepository
}

func New(repInterface absInterface.Repository) absInterface.Behavior {
	switch rep := repInterface.(type) {
	case interfaces.AuthRepository:
		return &Behavior{rep}
	default:
		return nil
	}
}

// RegisterNewUser - регистрация | добавление нового пользователя, генерация для него jwt
func (b *Behavior) RegisterNewUser(username string, password string) (bJWT.TokenString, *cErrors.MsgError) {

	// Проверка есть ли такой username
	// если произошла ошибка, вернуть её
	exists, errM := b.isUserExists(username)
	if exists == nil || *exists {
		return "", errM
	}

	// хэширование пароля
	hash, errM := hashPassword(password)
	if errM != nil {
		return "", errM
	}

	// сохранение юзера в БД и получение модельку пользователя
	user, errM := b.saveUser(username, hash)
	if errM != nil {
		return "", errM
	}

	// сгенерировать для пользователя токен
	token, err := createJWT(user)
	if err != nil {
		return "", err
	}

	// вернуть токен
	return token, nil
}

// AuthoriseUser - авторизация | проверяет существует ли пользователь, верный ли пароль, генерирует jwt для него
func (b *Behavior) AuthoriseUser(username string, password string) (bJWT.TokenString, *cErrors.MsgError) {
	op := "auth.api.AuthoriseUser"

	// проверяем существует ли пользователь
	exists, err := b.isUserExists(username)
	// если не существует или какая-то ошибка, выходим
	if exists == nil || !*exists {
		logger.Debug(fmt.Sprintf("Authorisation failed: user %s does not exist or err", username), op)
		return "", err
	}

	// получаем модельку юзера по username
	user, errM := b.getUser(username)
	if errM != nil {
		logger.Debug(fmt.Sprintf("Authorisation failed: user %s does not exist", username), op)
		return "", errM
	}

	// сравниваем пароли
	ok, errM := b.comparePassword(*user, password)
	if errM != nil || !ok {
		logger.Debug(fmt.Sprintf("Authorisation failed: user %s does not match", username), op)
		return "", errM
	}

	// сгенерировать для пользователя токен
	token, err := createJWT(user)
	if err != nil {
		logger.Debug(fmt.Sprintf("Authorisation failed: user %s generation token failed", username), op)
		return "", err
	}

	logger.StandardDebug(fmt.Sprintf("Login user={%v} with token={%v}", username, token), op)

	return token, nil
}

func (b *Behavior) isUserExists(username string) (*bool, *cErrors.MsgError) {
	op := "auth.behavior.IsUserExists"

	// Проверка есть ли такой username
	// если произошла ошибка, вернуть её
	exists, err := b.rep.UserExists(username)
	if err != nil {
		return nil, cErrors.UnknownError(err, op)
	}

	ok := true
	nOk := false
	if exists {
		return &ok, cErrors.New("user already exists", "пользователь уже существует")
	}
	return &nOk, cErrors.New("user doesn't exist", "пользователя не существует")
}

func hashPassword(password string) (string, *cErrors.MsgError) {
	op := "auth.behavior.HashPassword"

	hash, err := hasher.HashPassword(password)
	if err != nil {
		return "", cErrors.UnknownError(err, op)
	}
	return hash, nil
}

func (b *Behavior) saveUser(username string, hash string) (*bModels.User, *cErrors.MsgError) {
	op := "auth.behavior.SaveUser"

	role := bModels.Reader
	user, err := b.rep.SaveUser(username, int(role), hash)
	if err != nil {
		return nil, cErrors.UnknownError(err, op)
	}
	return user, nil
}

func createJWT(user *bModels.User) (bJWT.TokenString, *cErrors.MsgError) {
	op := "auth.behavior.CreateJWT"

	// сгенерировать для пользователя токен
	token, err := bJWT.CreateJWT(*user, global.TTL)
	if err != nil {
		return "", cErrors.UnknownError(err, op)
	}
	return token, nil
}

func (b *Behavior) getUser(username string) (*bModels.User, *cErrors.MsgError) {
	op := "auth.behavior.GetUser"

	user, err := b.rep.GetUserByUsername(username)
	if err != nil {
		return nil, cErrors.UnknownError(err, op)
	}
	return user, nil
}

func (b *Behavior) comparePassword(user bModels.User, password string) (bool, *cErrors.MsgError) {
	op := "auth.behavior.ComparePassword"

	userHash, err := b.rep.GetPasswordHashByID(user.UserID)
	if err != nil {
		return false, cErrors.UnknownError(err, op)
	}

	// Сравниваем введённый пароль с сохранённым хэшем
	err = hasher.CheckPasswordHash(password, userHash)
	if err != nil {
		return false, cErrors.New("hash mismatch", "некорректные данные")
	} else {
		return true, nil
	}
}
