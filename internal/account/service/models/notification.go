package models

// Service модель Notification
type Notification struct {
	// ID уведомления
	NotificationID string
	// Содержание
	Message string
	// ID отправителя
	SenderID string
	// Имя отправителя
	SenderUsername string
	// Статус: прочитано / нет
	IsRead bool
	// Дата отправления
	CreatedAt string
}
