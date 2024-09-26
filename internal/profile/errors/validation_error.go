package errors

// ValidationError - ошибка
// возникает при валидации полей
// нужна для сохранения и сообщения фронту, и внутренней ошибки
type ValidationError struct {
	err string
	msg string
}

func (m *ValidationError) Error() string {
	return m.err
}

// GetMsg - геттер для получение сообщения
func (m *ValidationError) GetMsg() string {
	return m.msg
}

// InitValidationError - инициализация ошибки
func NewValidationError(err string, msg string) *ValidationError {
	return &ValidationError{
		err: err,
		msg: msg,
	}
}
