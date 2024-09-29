package models

import (
	"fmt"
	er "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/errors"
)

// NewValidationFieldError - возвращает функцию для генерации ошибок
// при валидации в каждом поле.
// field string - название поле, в котором происходит валидация
// op string - путь до функции, в которой происходит ошибка
// формат op: service.dir_1.dir_2.MyValidateFunction
func NewValidationFieldError(field string, op string) func(msg string) *er.MsgError {
	return func(msg string) *er.MsgError {
		return er.New(
			fmt.Sprintf("%v is not valid msg={%v}| in %v", field, msg, op),
			msg)
	}
}
