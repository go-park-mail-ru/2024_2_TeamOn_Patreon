package models

import (
	"fmt"
	"regexp"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/pkg/errors"
)

// Service модель изменения данных аккаунта пользователя
type UpdateAccount struct {
	// Имя пользователя
	Username string `json:"username,omitempty"`
	// Пароль пользователя
	Password string `json:"password,omitempty"`
	// Старый пароль пользователя
	OldPassword string `json:"oldPassword,omitempty"`
	// Почта пользователя (если есть)
	Email string `json:"email,omitempty"`
	// Роль: читатель или автор
	Role string `json:"role"`
}

func (upd *UpdateAccount) String() string {
	return fmt.Sprintf("Update Account model {username: %s, password: %s, email: %v}", upd.Username, upd.Password, upd.Email)
}

func (upd *UpdateAccount) Validate() (bool, error) {
	op := "upd.Validate"

	err := upd.validateUsername()
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	err = upd.validatePassword()
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	return true, nil
}

func (upd *UpdateAccount) validateUsername() error {
	if upd.Username == "" {
		return nil
	}
	// Длина не менее 4 символов
	if len(upd.Username) < 4 {
		return global.ErrSmallLogin
	}

	// Длина не более 10 символов
	if len(upd.Username) > 10 {
		return global.ErrLongLogin
	}

	checks := []check{
		check{pattern: `^[a-zA-Z0-9_-]+$`,
			err: global.ErrLoginWithSpecChar},
		check{pattern: `|\s|`, err: global.ErrLoginWithSpace},
		check{pattern: `^[a-zA-Z]`, err: global.ErrLoginMustStartWithChar},
	}

	for _, chk := range checks {
		re := regexp.MustCompile(chk.pattern)
		if !re.MatchString(upd.Username) {
			return chk.err
		}
	}

	return nil
}

func (upd *UpdateAccount) validatePassword() error {
	if upd.Password == "" {
		return nil
	}
	// Длина не меньше 8 символов
	if len(upd.Password) < 8 {
		return global.ErrSmallPassword
	}

	// Длина не больше 64 символов
	if len(upd.Password) > 64 {
		return global.ErrLongPassword
	}

	checks := []check{
		check{pattern: `[!@#$%^&*()_+={}:|,.?]`, err: global.ErrPasswordWithoutSpecChar},
		check{pattern: `[A-Z]`, err: global.ErrPasswordWithoutUpperChar},
		check{pattern: `[a-z]`, err: global.ErrPasswordWithoutLowerChar},
		check{pattern: `[0-9]`, err: global.ErrPasswordWithoutNumberChar},
		check{pattern: `^[a-zA-Z0-9!@#$%^&*()_+={}:|,.?]+$`, err: global.ErrPasswordWithDifferentChar},
	}

	for _, chk := range checks {
		re := regexp.MustCompile(chk.pattern)
		if !re.MatchString(upd.Password) {
			return chk.err
		}
	}

	return nil
}
