package interfaces

import (
	"context"

	pkgModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/pkg/models"
	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/service/models"
)

// Интерфейс AccountService необходим для взаимодействия уровня controller с уровнем service
type AccountService interface {
	// GetAccDataByID - получение данных аккаунта по userID
	GetAccDataByID(ctx context.Context, userID string) (sModels.User, error)

	// GetAccSubscriptions - получение подписок аккаунта по userID
	GetAccSubscriptions(ctx context.Context, userID string) ([]sModels.Subscription, error)

	// GetAvatarByID - получение аватарки пользователя по userID
	GetAvatarByID(ctx context.Context, userID string) ([]byte, error)

	// UpdateUsername - изменение имени аккаунта по userID
	UpdateUsername(ctx context.Context, userID string, username string) error

	// UpdatePassword - изменение пароля аккаунта по userID
	UpdatePassword(ctx context.Context, userID, oldPassword, password string) error

	// UpdateEmail - изменение почты аккаунта по userID
	UpdateEmail(ctx context.Context, userID string, email string) error

	// PostAccountUpdateAvatar - изменение аватарки аккаунта по userID
	PostUpdateAvatar(ctx context.Context, userID string, file []byte, fileExtension string) error

	// PostUpdateRole - изменение аватарки аккаунта по userID
	PostUpdateRole(ctx context.Context, userID string) error

	// NOTIFICATIONS

	// GetNotifications - получение уведомлений пользователя
	GetNotifications(ctx context.Context, userID string, opt *pkgModels.NotificationsOpt) ([]*pkgModels.Notification, error)

	// GetNewNotifications - получение последних уведомлений пользователя
	GetNewNotifications(ctx context.Context, userID string, opt *pkgModels.NotificationsTimeOpt) ([]*pkgModels.Notification, error)

	// ReadNotification - изменяет статус уведомления на "прочитано"
	ReadNotification(ctx context.Context, userID, notificationID string) error
}
