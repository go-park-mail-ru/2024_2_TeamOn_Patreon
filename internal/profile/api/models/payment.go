package models

// Модель выплаты
type Payment struct {
	// Сумма выплаты
	Amount float64 `json:"amount"`
}
