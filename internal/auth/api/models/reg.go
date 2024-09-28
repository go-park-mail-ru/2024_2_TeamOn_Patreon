package models

import (
	"regexp"

	er "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/errors"
)

// Reg - модель для фронта
type Reg struct {
	// Имя пользователя. допустимые символы - латинские буквы, цифры и \"-\", \"_\".
	Username string `json:"username"`
	// Пароль. должен содержать хотя бы 1 заглавную, 1 строчную латинские буквы, 1 цифру, 1 спец символ.
	Password string `json:"password"`
}

func (reg *Reg) Validate() (bool, *er.MsgError) {
	userValid, err := reg.validateUsername()
	if err != nil || !userValid {
		return false, err
	}

	passValid, err := reg.validatePassword()
	if err != nil || !passValid {
		return false, err
	}

	return true, nil
}

func (reg *Reg) validateUsername() (bool, *er.MsgError) {
	op := "auth.api.reg.validateUsername"
	validErr := InitValidErrorInField("username", op)

	// Длина не менее 4 символов
	if len(reg.Username) < 4 {
		msg := "логин должен быть не меньше 4 символов"
		return false, validErr(msg)
	}

	// Длина не более 10 символов
	if len(reg.Username) > 10 {
		msg := "логин должен быть не больше 10 символов"
		return false, validErr(msg)
	}

	checks := []check{
		check{pattern: `^[a-zA-Z0-9_-]+$`,
			msg: "логин должен содержать только латинские символы, цифры и символы '-', '_'"},
		check{pattern: `|\s|`, msg: "логин не должен содержать пробелы"},
		check{pattern: `^[a-zA-Z]`, msg: "логин должен начинаться с буквы"},
	}

	for _, chk := range checks {
		re := regexp.MustCompile(chk.pattern)
		if !re.MatchString(reg.Username) {
			return false, validErr(chk.msg)
		}
	}

	return true, nil
}

func (reg *Reg) validatePassword() (bool, *er.MsgError) {
	op := "auth.api.reg.validatePassword"
	validErr := InitValidErrorInField("password", op)

	// Длина не меньше 8 символов
	if len(reg.Password) < 8 {
		msg := "пароль должен быть не меньше 8 символов"
		return false, validErr(msg)
	}

	// Длина не больше 64 символов
	if len(reg.Password) > 64 {
		msg := "пароль должен быть не больше 64 символов"
		return false, validErr(msg)
	}

	checks := []check{
		check{pattern: `[!@#$%^&*()_+={}:|,.?]`, msg: "пароль должен содержать спец символ"},
		check{pattern: `[A-Z]`, msg: "пароль должен содержать латинскую букву в верхнем регистре"},
		check{pattern: `[a-z]`, msg: "пароль должен содержать латинскую букву в нижнем регистре"},
		check{pattern: `[0-9]`, msg: "пароль должен содержать цифры"},
		check{pattern: `^[a-zA-Z0-9!@#$%^&*()_+={}:|,.?]+$`, msg: "пароль может содержать только буквы, цифры и спец символы: ^[a-zA-Z0-9!@#$%^&*()_+={}:|,.?]+$"},
	}

	for _, chk := range checks {
		re := regexp.MustCompile(chk.pattern)
		if !re.MatchString(reg.Password) {
			return false, validErr(chk.msg)
		}
	}

	return true, nil
}
