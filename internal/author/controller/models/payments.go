package models

//go:generate easyjson

// Модель выплат
//
//easyjson:json
type Payments struct {
	// Сумма выплат
	Amount int `json:"amount"`
}
