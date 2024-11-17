package models

// SubscriptionLayer Данные об уровне подписки, которые возвращает бэк
type SubscriptionLayer struct {
	// Уровень подписки. Уровень [0:3]
	Layer int `json:"layer,omitempty"`
	// Название этого уровня по умолчанию
	LayerName string `json:"layerName,omitempty"`
}
