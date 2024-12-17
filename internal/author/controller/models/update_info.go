package models

//go:generate easyjson -all

// Service модель обновления информации об авторе
//
//easyjson:json
type UpdateInfo struct {
	// Имя пользователя
	Info string `json:"info"`
}
