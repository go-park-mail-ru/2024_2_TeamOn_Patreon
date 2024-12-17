package service

import (
	"context"

	hasher "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service/hasher"
	rInterfaces "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service/interfaces"
	bJWT "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service/jwt"
	global "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

type Behavior struct {
	rep rInterfaces.AuthRepository
}

func New(repository rInterfaces.AuthRepository) *Behavior {
	return &Behavior{repository}
}

// AuthoriseUser - авторизация | проверяет существует ли пользователь, верный ли пароль, генерирует jwt для него
func (b *Behavior) LoginUser(ctx context.Context, username string, password string) (bJWT.TokenString, error) {
	op := "internal.auth.behavior.LoginUser"

	// проверяем существует ли пользователь
	exists, err := b.isUserExists(ctx, username)
	// если не существует или какая-то ошибка, выходим
	if err != nil {
		logger.StandardDebugF(ctx, op, "Authorisation failed: user %s does not exist or err", username)
		return "", errors.Wrap(err, op)
	}
	if !exists {
		logger.StandardDebugF(ctx, op, "Authorisation failed: user %s does not exist", username)
		return "", global.ErrNotValidUserAndPassword
	}

	logger.StandardDebugF(ctx, op, "Getting user of username %s", username)
	// получаем модельку юзера по username
	user, err := b.getUser(ctx, username)
	if err != nil {
		logger.StandardDebugF(ctx, op, "Authorisation failed: user %s does not exist in DB", username)
		return "", errors.Wrap(err, op)
	}

	logger.StandardDebugF(ctx, op, "Getting user=%v", user)

	// сравниваем пароли
	ok, err := b.comparePassword(ctx, *user, password)
	if err != nil {
		logger.StandardDebugF(ctx, op, "Authorisation failed: user %s does not match", username)
		return "", errors.Wrap(err, op)
	}
	if !ok {
		logger.StandardDebugF(ctx, op, "Authorisation failed: user %s does not match", username)
		return "", errors.Wrap(global.ErrNotValidUserAndPassword, op)
	}

	// сгенерировать для пользователя токен
	token, err := createJWT(user)
	if err != nil {
		logger.StandardDebugF(ctx, op, "Authorisation failed: user %s generation token failed", username)
		return "", errors.Wrap(err, op)
	}

	logger.StandardDebugF(ctx, op, "Login user={%v} with token={%v}", username, token)

	return token, nil
}

func (b *Behavior) isUserExists(ctx context.Context, username string) (bool, error) {
	op := "auth.service.IsUserExists"

	// Проверка есть ли такой username
	// если произошла ошибка, вернуть её
	exists, err := b.rep.UserExists(ctx, username)
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	return exists, nil
}

func hashPassword(password string) (string, error) {
	op := "auth.service.HashPassword"

	hash, err := hasher.HashPassword(password)
	if err != nil {
		return "", errors.Wrap(global.ErrServer, op)
	}
	return hash, nil
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

func (b *Behavior) getUser(ctx context.Context, username string) (*bModels.User, error) {
	op := "auth.service.GetUser"

	userId, role, err := b.rep.GetUserByUsername(ctx, username)
	logger.StandardDebugF(ctx, op, "Got userId=%v role=%v by username=%v", userId, role, username)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	user := &bModels.User{UserID: bModels.UserID(userId.String()), Role: bModels.Role(role), Username: username}
	return user, nil
}

func (b *Behavior) comparePassword(ctx context.Context, user bModels.User, password string) (bool, error) {
	op := "auth.service.ComparePassword"

	userIdUuid, err := uuid.FromString(string(user.UserID))
	if err != nil {
		return false, errors.Wrap(err, op)
	}
	userHash, err := b.rep.GetPasswordHashByID(ctx, userIdUuid)
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	// Сравниваем введённый пароль с сохранённым хэшем
	err = hasher.CheckPasswordHash(password, userHash)
	if err != nil {
		return false, errors.Wrap(global.ErrNotValidUserAndPassword, op)
	} else {
		return true, nil
	}
}

func (b *Behavior) LogoutUser(ctx context.Context, userID bModels.UserID) error {
	op := "auth.service.behavior.LogoutUser"

	logger.StandardInfoF(ctx, op, "LogoutUser user={%v}", userID)
	return nil
}

func (b *Behavior) generateUUID() uuid.UUID {
	return uuid.NewV4()
}
