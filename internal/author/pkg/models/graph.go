package models

//go:generate easyjson

//easyjson:json
type Graphic struct {
	PointsX []int `json:"valueX"`
	PointsY []int `json:"valueY"`
}
