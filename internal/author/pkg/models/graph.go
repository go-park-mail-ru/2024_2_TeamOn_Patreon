package models

//go:generate easyjson -all

//easyjson:json
type Graphic struct {
	PointsX []int `json:"valueX"`
	PointsY []int `json:"valueY"`
}
