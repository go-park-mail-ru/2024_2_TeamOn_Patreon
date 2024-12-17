package models

import (
	"context"
	"fmt"
	"regexp"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/validate"
	"github.com/pkg/errors"
)

//go:generate easyjson -all

//easyjson:json
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
	op := "auth.controller.login.Validate"

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
	op := "auth.controller.login.validateUsername"

	lg.Username = validate.Sanitize(lg.Username)

	// Длина не менее 4 символов
	if len(lg.Username) < 4 {
		logger.StandardDebugF(context.Background(), op, "login isn't valid: '%v'", lg.Username)
		return global.ErrSmallLogin
	}

	// Длина не более 10 символов
	if len(lg.Username) > 10 {
		return global.ErrLongLogin
	}

	checks := []check{
		//check{pattern: `^[a-zA-Z0-9_-]+$`, err: global.ErrLoginWithSpecChar},
		check{pattern: `|\s|`, err: global.ErrLoginWithSpace},
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
	//op := "auth.controller.login.validatePassword"

	lg.Password = validate.Sanitize(lg.Password)

	// Длина не меньше 8 символов
	if len(lg.Password) < 8 {
		return global.ErrSmallPassword
	}

	// Длина не больше 64 символов
	if len(lg.Password) > 64 {
		return global.ErrLongPassword
	}

	checks := []check{
		check{pattern: `^[a-zA-Z0-9!@#$%^&*()_+={}:|,.?]+$`, err: global.ErrPasswordWithDifferentChar},
	}

	for _, chk := range checks {
		re := regexp.MustCompile(chk.pattern)
		if !re.MatchString(lg.Password) {
			return chk.err
		}
	}

	return nil
}
