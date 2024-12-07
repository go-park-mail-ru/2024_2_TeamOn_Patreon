package models

import (
	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/service/models"
	valid "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/validate"
)

// Controller модель Notification
type Notification struct {
	// ID уведомления
	NotificationID string `json:"notificationID"`
	// Содержание
	Message string `json:"message"`
	// ID отправителя
	SenderID string `json:"senderID"`
	// Имя отправителя
	SenderUsername string `json:"senderUsername"`
	// Статус: прочитано / нет
	IsRead bool `json:"isRead"`
	// Дата отправления
	CreatedAt string `json:"createdAt"`
}

func mapNotifcationServiceToController(ntf sModels.Notification) *Notification {
	return &Notification{
		NotificationID: valid.Sanitize(ntf.NotificationID),
		Message:        valid.Sanitize(ntf.Message),
		SenderID:       valid.Sanitize(ntf.SenderID),
		SenderUsername: valid.Sanitize(ntf.SenderUsername),
		IsRead:         ntf.IsRead,
		CreatedAt:      valid.Sanitize(ntf.CreatedAt),
	}
}

func MapNotifcationsServToCntrl(sNotifications []*sModels.Notification) []*Notification {
	notifications := make([]*Notification, 0, len(sNotifications))
	for _, ntf := range sNotifications {
		notifications = append(notifications, mapNotifcationServiceToController(*ntf))
	}
	return notifications
}
