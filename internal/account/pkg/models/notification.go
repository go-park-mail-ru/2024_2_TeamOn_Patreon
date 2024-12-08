package models

import "time"

// Repository модель Notification
type Notification struct {
	// ID уведомления
	NotificationID string
	// Содержание
	Message string
	// ID отправителя
	SenderID string
	// Статус: прочитано / нет
	IsRead bool
	// Дата отправления
	CreatedAt time.Time
}
