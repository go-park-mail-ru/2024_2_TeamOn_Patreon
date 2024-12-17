package models

//go:generate easyjson

// Модель выплаты
//
//easyjson:json
type Payment struct {
	// Сумма выплаты
	Amount float64 `json:"amount"`
}
