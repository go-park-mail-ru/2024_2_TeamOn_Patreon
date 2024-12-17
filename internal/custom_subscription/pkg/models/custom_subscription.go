package models

import (
	"fmt"
)

type SubscriptionLayer struct {
	Layer     int
	LayerName string
}

func (sl *SubscriptionLayer) String() string {
	return fmt.Sprintf("Layer: %d, LayerName: %s", sl.Layer, sl.LayerName)
}

type CustomSubscription struct {
	CustomSubscriptionID string
	Title                string
	Description          string
	Cost                 int
	Layer                int
}

func (sc CustomSubscription) String() string {
	return fmt.Sprintf("Custom Subscription ID: %v, Title: %s, Description: %s, Cost: %d, Layer: %d",
		sc.CustomSubscriptionID, sc.Title, sc.Description, sc.Cost, sc.Layer)
}
