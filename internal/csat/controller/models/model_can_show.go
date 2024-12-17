package models

//go:generate easyjson

//easyjson:json
type ModelCanShow struct {
	// Флаг, можно ли задать юзеру вопрос
	CanAsk bool `json:"isCanShow"`
}
