package models

import (
	"time"

	pkgModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/pkg/models"
	valid "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/validate"
)

// Controller модель NotificationID
type NotificationID struct {
	// ID уведомления
	ID string `json:"notificationID"`
}

// Controller модель Notification
type Notification struct {
	// ID уведомления
	NotificationID string `json:"notificationID"`
	// Содержание
	Message string `json:"message"`
	// ID отправителя
	SenderID string `json:"senderID"`
	// Статус: прочитано / нет
	IsRead bool `json:"isRead"`
	// Дата отправления
	CreatedAt time.Time `json:"createdAt"`
}

func mapNotificationCommonToController(ntf pkgModels.Notification) *Notification {
	return &Notification{
		NotificationID: valid.Sanitize(ntf.NotificationID),
		Message:        valid.Sanitize(ntf.Message),
		SenderID:       valid.Sanitize(ntf.SenderID),
		IsRead:         ntf.IsRead,
		CreatedAt:      ntf.CreatedAt,
	}
}

func MapNotificationsCommonToController(sNotifications []*pkgModels.Notification) []*Notification {
	notifications := make([]*Notification, 0, len(sNotifications))
	for _, ntf := range sNotifications {
		notifications = append(notifications, mapNotificationCommonToController(*ntf))
	}
	return notifications
}
