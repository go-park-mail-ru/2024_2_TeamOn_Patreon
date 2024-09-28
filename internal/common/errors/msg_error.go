package errors

// MsgError - ошибка
// нужна для сохранения и сообщения фронту, и внутренней ошибки
type MsgError struct {
	err string // Текст внутренней ошибки
	msg string // Текст для сообщений фронту
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
