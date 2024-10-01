package models

import (
	"fmt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/config"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/logger"
	"github.com/pkg/errors"
	"regexp"
)

type Login struct {
	// Логин пользователя (имя пользователя или почта)
	Username string `json:"username"`
	// Пароль пользователя
	Password string `json:"password"`
}

func (lg *Login) String() string {
	return fmt.Sprintf("Login model {Username: %s, Password: %s}", lg.Username, lg.Password)
}

func (lg *Login) Validate() (bool, error) {
	op := "auth.api.login.Validate"

	err := lg.validateUsername()
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	err = lg.validatePassword()
	if err != nil {
		return false, errors.Wrap(err, op)
	}
	return true, nil
}

func (lg *Login) validateUsername() error {
	op := "auth.api.login.validateUsername"

	// Длина не менее 4 символов
	if len(lg.Username) < 4 {
		logger.StandardDebugF(op, "login isn't valid: '%v'", lg.Username)
		return config.ErrSmallLogin
	}

	// Длина не более 10 символов
	if len(lg.Username) > 10 {
		return config.ErrLongLogin
	}

	checks := []check{
		check{pattern: `^[a-zA-Z0-9_-]+$`, err: config.ErrLoginWithSpecChar},
		check{pattern: `|\s|`, err: config.ErrLoginWithSpace},
	}

	for _, chk := range checks {
		re := regexp.MustCompile(chk.pattern)
		if !re.MatchString(lg.Username) {
			return chk.err
		}
	}

	return nil
}

func (lg *Login) validatePassword() error {
	//op := "auth.api.login.validatePassword"

	// Длина не меньше 8 символов
	if len(lg.Password) < 8 {
		return config.ErrSmallPassword
	}

	// Длина не больше 64 символов
	if len(lg.Password) > 64 {
		return config.ErrLongPassword
	}

	checks := []check{
		check{pattern: `^[a-zA-Z0-9!@#$%^&*()_+={}:|,.?]+$`, err: config.ErrPasswordWithDifferentChar},
	}

	for _, chk := range checks {
		re := regexp.MustCompile(chk.pattern)
		if !re.MatchString(lg.Password) {
			return chk.err
		}
	}

	return nil
}
