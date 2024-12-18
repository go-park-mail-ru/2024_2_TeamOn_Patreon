package models

//go:generate easyjson

//easyjson:json
type RatingModel struct {
	Rating int `json:"rating"`
}
