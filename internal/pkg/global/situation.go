package global

import (
	"github.com/pkg/errors"
	"net/http"
)

// ВРЕМЕННО В ЭТОМ ФАЙЛЕ ПОТОМ РАЗНЕСЕМ
// ТАКЖЕ КАК И ТРЕБОВАНИЯ ПО ВАЛИДАЦИИ

var (
	ErrBadRequest = errors.New("bad request")
	// VALIDATION
	// registration

	ErrSmallLogin             = errors.New("login is too short")
	ErrLongLogin              = errors.New("login is too long")
	ErrLoginWithSpecChar      = errors.New("login with spec char")
	ErrLoginWithSpace         = errors.New("login with spec space")
	ErrLoginMustStartWithChar = errors.New("login must start with char")

	ErrSmallPassword             = errors.New("password is too short")
	ErrLongPassword              = errors.New("password is too long")
	ErrPasswordWithoutSpecChar   = errors.New("password without spec char")
	ErrPasswordWithoutLowerChar  = errors.New("password without lower char")
	ErrPasswordWithoutUpperChar  = errors.New("password without upper char")
	ErrPasswordWithoutNumberChar = errors.New("password without number char")
	ErrPasswordWithDifferentChar = errors.New("password with different char")

	// BEHAVIOR

	// register

	ErrUserAlreadyExists = errors.New("user already exists")

	// auth

	ErrUserNotFound            = errors.New("user not found")
	ErrNotValidUserAndPassword = errors.New("not valid user and password")
	ErrNotEnoughRights         = errors.New("not enough rights")

	// logout

	ErrUserNotAuthorized = errors.New("user not authorized")

	// database

	ErrDatabaseDead = errors.New("database is not responding")
	ErrServer       = errors.New("end-to-end error")

	// json is incorrect

	ErrInvalidJSON = errors.New("invalid JSON format")

	// content
	// validate

	ErrFieldTooLong             = errors.New("field too long")
	ErrFieldTooShort            = errors.New("field too short")
	ErrFieldContainsSpecialChar = errors.New("field contains special char")

	// post

	ErrPostDoesntExists = errors.New("post doesn't")

	// uuid

	ErrUuidIsInvalid = errors.New("uuid is invalid")
)

type ErrorHttpInfo struct {
	msg  string
	code int
}

// ВРЕМЕННО ЗДЕСЬ
var mapErrToHttpModel = map[error]ErrorHttpInfo{
	ErrBadRequest: {msg: "bad request", code: http.StatusBadRequest},

	ErrSmallLogin: {msg: "логин должен быть не меньше 4 символов", code: http.StatusBadRequest},
	ErrLongLogin:  {msg: "логин должен быть не более 10 символов", code: http.StatusBadRequest},

	ErrLoginWithSpecChar: {msg: "логин должен содержать только латинские символы, цифры и символы '-', '_'",
		code: http.StatusBadRequest},
	ErrLoginWithSpace:         {msg: "логин не должен содержать пробелы", code: http.StatusBadRequest},
	ErrLoginMustStartWithChar: {msg: "логин должен начинаться с буквы", code: http.StatusBadRequest},

	ErrSmallPassword: {msg: "пароль должен быть не меньше 8 символов", code: http.StatusBadRequest},
	ErrLongPassword:  {msg: "пароль должен быть не больше 64 символов", code: http.StatusBadRequest},

	ErrPasswordWithoutSpecChar: {msg: "пароль должен содержать спец символ", code: http.StatusBadRequest},
	ErrPasswordWithoutLowerChar: {msg: "пароль должен содержать латинскую букву в нижнем регистре",
		code: http.StatusBadRequest},
	ErrPasswordWithoutUpperChar:  {msg: "пароль должен содержать латинскую букву в верхнем регистре", code: http.StatusBadRequest},
	ErrPasswordWithoutNumberChar: {msg: "пароль должен содержать цифры", code: http.StatusBadRequest},
	ErrPasswordWithDifferentChar: {msg: "пароль может содержать только буквы, цифры и спец символы", code: http.StatusBadRequest},

	// AUTH
	ErrUserNotFound:            {msg: "некорректные данные", code: http.StatusBadRequest},
	ErrNotValidUserAndPassword: {msg: "некорректные данные", code: http.StatusBadRequest},
	ErrUserAlreadyExists:       {msg: "пользователь уже существует", code: http.StatusBadRequest},
	ErrInvalidJSON:             {msg: "невалидный запрос", code: http.StatusBadRequest},
	ErrDatabaseDead:            {msg: "ошибка сервера", code: http.StatusInternalServerError},

	ErrServer:            {msg: "end-to-end error", code: http.StatusInternalServerError},
	ErrUserNotAuthorized: {msg: "пользователь не авторизован", code: http.StatusUnauthorized},

	// content
	ErrFieldTooLong:             {msg: "поле слишком длинное", code: http.StatusBadRequest},
	ErrFieldTooShort:            {msg: "поле слишком короткое", code: http.StatusBadRequest},
	ErrFieldContainsSpecialChar: {msg: "поле содержит запрещенные символы", code: http.StatusBadRequest},

	// uuid
	ErrUuidIsInvalid: {msg: "невалидный uuid", code: http.StatusBadRequest},

	// rights
	ErrNotEnoughRights:  {msg: "недостаточно прав", code: http.StatusBadRequest},
	ErrPostDoesntExists: {msg: "пост не найден", code: http.StatusNoContent},
}

func GetMsgError(err error) string {
	err = errors.Cause(err)
	if httpInfo, exists := mapErrToHttpModel[err]; exists {
		return httpInfo.msg
	}
	// Если ошибка не найдена в мапе, возвращаем общее сообщение
	return "некая ошибка сервера"
}

func GetCodeError(err error) int {
	err = errors.Cause(err)
	if httpInfo, exists := mapErrToHttpModel[err]; exists {
		return httpInfo.code
	}
	// Если ошибка не найдена в мапе, возвращаем общее сообщение
	return http.StatusNotImplemented
}
