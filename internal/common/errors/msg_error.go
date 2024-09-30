package errors

import "net/http"

// MsgError - ошибка
// нужна для сохранения и сообщения фронту, и внутренней ошибки
type MsgError struct {
	err  string // Текст внутренней ошибки
	msg  string // Текст для сообщений фронту
	code int    // http код ошибки
}

// Error возвращает ошибку
func (m *MsgError) Error() string {
	return m.err
}

// GetMsg - геттер для получения сообщения
func (m *MsgError) GetMsg() string {
	return m.msg
}

// New - инициализация ошибки
func New(err string, msg string) *MsgError {
	return &MsgError{
		err: err,
		msg: msg,
	}
}

func NewCode(err string, msg string, code int) *MsgError {
	return &MsgError{
		err:  err,
		msg:  msg,
		code: code,
	}
}

func (m *MsgError) Code() int {
	if m.code == 0 {
		return http.StatusInternalServerError
	}
	return m.code
}
