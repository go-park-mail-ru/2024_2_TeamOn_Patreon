package models

import (
	"fmt"
	er "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/errors"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/logger"
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

func (lg *Login) Validate() (bool, *er.MsgError) {
	userValid, err := lg.validateUsername()
	if err != nil || !userValid {
		return false, err
	}

	passValid, err := lg.validatePassword()
	if err != nil || !passValid {
		return false, err
	}
	return true, nil
}

func (lg *Login) validateUsername() (bool, *er.MsgError) {
	op := "auth.api.login.validateUsername"
	validErr := InitValidErrorInField("username", op)

	// Длина не менее 4 символов
	if len(lg.Username) < 4 {
		msg := "логин должен быть не меньше 4 символов"
		logger.StandardDebugF(op, "login isn't valid: '%v'", lg.Username)
		return false, validErr(msg)
	}

	// Длина не более 10 символов
	if len(lg.Username) > 10 {
		msg := "логин должен быть не больше 10 символов"
		return false, validErr(msg)
	}

	checks := []check{
		check{pattern: `^[a-zA-Z0-9_-]+$`,
			msg: "логин должен содержать только латинские символы, цифры и символы '-', '_'"},
		check{pattern: `|\s|`, msg: "логин не должен содержать пробелы"},
	}

	for _, chk := range checks {
		re := regexp.MustCompile(chk.pattern)
		if !re.MatchString(lg.Username) {
			return false, validErr(chk.msg)
		}
	}

	return true, nil
}

func (lg *Login) validatePassword() (bool, *er.MsgError) {
	op := "auth.api.login.validatePassword"
	validErr := InitValidErrorInField("password", op)

	// Длина не меньше 8 символов
	if len(lg.Password) < 8 {
		msg := "пароль должен быть не меньше 8 символов"
		return false, validErr(msg)
	}

	// Длина не больше 64 символов
	if len(lg.Password) > 64 {
		msg := "пароль должен быть не больше 64 символов"
		return false, validErr(msg)
	}

	checks := []check{
		check{pattern: `^[a-zA-Z0-9!@#$%^&*()_+={}:|,.?]+$`,
			msg: "пароль может содержать только буквы, цифры и спец символы: ^[a-zA-Z0-9!@#$%^&*()_+={}:|,.?]+$"},
	}

	for _, chk := range checks {
		re := regexp.MustCompile(chk.pattern)
		if !re.MatchString(lg.Password) {
			return false, validErr(chk.msg)
		}
	}

	return true, nil
}
