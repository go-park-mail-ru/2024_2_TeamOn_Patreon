package models

import (
	"fmt"
	er "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/polina-auth/internal/common/errors"
)

// InitValidErrorInField - возвращает функцию для генерации ошибок
// при валидации в каждом поле.
// field string - название поле, в котором происходит валидация
// op string - путь до функции, в которой происходит ошибка
// формат op: service.dir_1.dir_2.MyValidateFunction
func InitValidErrorInField(field string, op string) func(msg string) *er.MsgError {
	return func(msg string) *er.MsgError {
		return er.New(
			fmt.Sprintf("%v is not valid msg={%v}| in %v", field, msg, op),
			msg)
	}
}
