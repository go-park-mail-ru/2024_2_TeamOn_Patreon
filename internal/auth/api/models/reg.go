package models

import (
	"fmt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/config"
	"github.com/pkg/errors"
	"regexp"
)

// Reg - модель для фронта
type Reg struct {
	// Имя пользователя. допустимые символы - латинские буквы, цифры и \"-\", \"_\".
	Username string `json:"username"`
	// Пароль. должен содержать хотя бы 1 заглавную, 1 строчную латинские буквы, 1 цифру, 1 спец символ.
	Password string `json:"password"`
}

func (reg *Reg) String() string {
	return fmt.Sprintf("Reg model {username: %s, password: %s}", reg.Username, reg.Password)
}

func (reg *Reg) Validate() (bool, error) {
	op := "reg.Validate"

	err := reg.validateUsername()
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	err = reg.validatePassword()
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	return true, nil
}

func (reg *Reg) validateUsername() error {
	// Длина не менее 4 символов
	if len(reg.Username) < 4 {
		return config.ErrSmallLogin
	}

	// Длина не более 10 символов
	if len(reg.Username) > 10 {
		return config.ErrLongLogin
	}

	checks := []check{
		check{pattern: `^[a-zA-Z0-9_-]+$`,
			err: config.ErrLoginWithSpecChar},
		check{pattern: `|\s|`, err: config.ErrLoginWithSpace},
		check{pattern: `^[a-zA-Z]`, err: config.ErrLoginMustStartWithChar},
	}

	for _, chk := range checks {
		re := regexp.MustCompile(chk.pattern)
		if !re.MatchString(reg.Username) {
			return chk.err
		}
	}

	return nil
}

func (reg *Reg) validatePassword() error {

	// Длина не меньше 8 символов
	if len(reg.Password) < 8 {
		return config.ErrSmallPassword
	}

	// Длина не больше 64 символов
	if len(reg.Password) > 64 {
		return config.ErrLongPassword
	}

	checks := []check{
		check{pattern: `[!@#$%^&*()_+={}:|,.?]`, err: config.ErrPasswordWithoutSpecChar},
		check{pattern: `[A-Z]`, err: config.ErrPasswordWithoutUpperChar},
		check{pattern: `[a-z]`, err: config.ErrPasswordWithoutLowerChar},
		check{pattern: `[0-9]`, err: config.ErrPasswordWithoutNumberChar},
		check{pattern: `^[a-zA-Z0-9!@#$%&*()_+=:,.?]+$`, err: config.ErrPasswordWithDifferentChar},
	}

	for _, chk := range checks {
		re := regexp.MustCompile(chk.pattern)
		if !re.MatchString(reg.Password) {
			return chk.err
		}
	}

	return nil
}
