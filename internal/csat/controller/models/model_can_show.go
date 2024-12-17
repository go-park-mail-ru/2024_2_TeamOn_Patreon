package models

//go:generate easyjson -all

//easyjson:json
type ModelCanShow struct {
	// Флаг, можно ли задать юзеру вопрос
	CanAsk bool `json:"isCanShow"`
}
