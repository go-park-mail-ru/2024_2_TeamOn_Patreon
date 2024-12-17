package models

import "fmt"

//go:generate easyjson

// SubscriptionLayer Данные об уровне подписки, которые возвращает бэк
//
//easyjson:json
type SubscriptionLayer struct {
	// Уровень подписки. Уровень [0:3]
	Layer int `json:"layer"`
	// Название этого уровня по умолчанию
	LayerName string `json:"layerName"`
}

func (sl *SubscriptionLayer) String() string {
	return fmt.Sprintf("Layer: %d, LayerName: %s", sl.Layer, sl.LayerName)
}
