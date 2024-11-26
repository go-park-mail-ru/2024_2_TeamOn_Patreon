package models

type ModelCanShow struct {
	// Флаг, можно ли задать юзеру вопрос
	CanAsk bool `json:"isCanShow"`
}
