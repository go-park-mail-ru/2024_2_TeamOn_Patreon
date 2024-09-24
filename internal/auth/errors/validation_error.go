/*
** HOW TO USE ?
pls read me

1. Импортируем модуль. Например, на данный момент, импорт такой:
```
import er "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/polina-auth/internal/auth/errors"
```

2. В функции валидации поля создаем функцию для создания ошибки именно для этого поля
Например:
```
op := "service.dir_1.dir_2.SomeFunctionValidateField"
validErr := er.InitValidErrorInField("my_field", op)
```

3. Везде где в этой функции валидации поля нужно возвращать ошибку
Например:
```
// Длина не более 10 символов
if len(reg.Username) > 10 {
	msg := "логин должен быть не больше 10 символов"
	return false, validErr(msg)
}
```

4. При отправки сообщения фронту - используем метод GetMsg
Например:
```
msgToFront :=  errV.GetMsg()
// отправка фронту
```

____________________________
P.S. Также поддерживается функционал:
- InitValidationError
	инициализация ошибки только с инициализацией полей err и msg
*/

package errors

import "fmt"

// ValidationError - ошибка
// возникает при валидации полей
// нужна для сохранения и сообщения фронту, и внутренней ошибки
type ValidationError struct {
	err string // текст внутренней ошибки
	msg string // текст для сообщений фронту
}

func (m *ValidationError) Error() string {
	return m.err
}

// GetMsg - геттер для получение сообщения
func (m *ValidationError) GetMsg() string {
	return m.msg
}

// InitValidationError - инициализация ошибки
func InitValidationError(err string, msg string) *ValidationError {
	return &ValidationError{
		err: err,
		msg: msg,
	}
}

// InitValidErrorInField - возвращает функцию для генерации ошибок
// при валидации в каждом поле.
// field string - название поле, в котором происходит валидация
// op string - путь до функции, в которой происходит ошибка
// формат op: service.dir_1.dir_2.MyValidateFunction
// (не уверена как назвать)
// ((извините за замыкания, практика нужна же))
func InitValidErrorInField(field string, op string) func(msg string) *ValidationError {
	return func(msg string) *ValidationError {
		return InitValidationError(
			fmt.Sprintf("%v is not valid msg={%v}| in %v", field, msg, op),
			msg)
	}
}
