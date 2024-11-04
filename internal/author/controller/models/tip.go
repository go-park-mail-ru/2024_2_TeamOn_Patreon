package models

// Controller модель пожертвования
type Tip struct {
	// Сообщение от донатера
	Message string `json:"message,omitempty"`
	// Сумма пожертвования
	Cost int `json:"cost,omitempty"`
}
