package models

//go:generate easyjson -all

//easyjson:json
type RatingModel struct {
	Rating int `json:"rating"`
}
