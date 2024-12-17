package models

//go:generate easyjson -all

// Модель выплат
//
//easyjson:json
type Payments struct {
	// Сумма выплат
	Amount int `json:"amount"`
}
