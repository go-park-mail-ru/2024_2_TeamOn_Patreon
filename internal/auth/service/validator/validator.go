package validator

import (
	"context"
	"regexp"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/pkg"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
)

func validateUsername(username string) (string, error) {
	op := "internal.auth.service.validator.validateUsername"
	// Длина не менее 4 символов
	if len(username) < pkg.MinUsernameLen {
		logger.StandardDebugF(context.Background(), op, "Username is too short username=%v", username)
		return username, errors.Wrap(global.ErrNotValidUserAndPassword, op)
	}

	// Длина не более 10 символов
	if len(username) > pkg.MaxUsernameLen {
		return username, errors.Wrap(global.ErrNotValidUserAndPassword, op)
	}

	return username, nil
}

func validatePassword(password string) (string, error) {
	op := "internal.auth.service.validator.validatePassword"

	// Длина не меньше 8 символов
	if len(password) < pkg.MinPasswordLen {
		return password, errors.Wrap(global.ErrNotValidUserAndPassword, op)
	}

	// Длина не больше 64 символов
	if len(password) > pkg.MaxPasswordLen {
		return password, errors.Wrap(global.ErrNotValidUserAndPassword, op)
	}

	re := regexp.MustCompile(pkg.PatternPasswordMustContainDiffChar)
	if !re.MatchString(password) {
		return password, errors.Wrap(global.ErrNotValidUserAndPassword, op)
	}

	return password, nil
}

func ValidateUsernameAndPassword(username string, password string) (string, string, error) {
	op := "internal.auth.service.validator.ValidateUsernameAndPassword"

	username, err := validateUsername(username)
	if err != nil {
		return username, password, errors.Wrap(err, op)
	}

	password, err = validatePassword(password)
	if err != nil {
		return username, password, errors.Wrap(err, op)
	}
	return username, password, nil
}
