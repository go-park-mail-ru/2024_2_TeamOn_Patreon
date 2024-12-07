package interfaces

import (
	"context"

	pkgModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/pkg/models"
	repModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/repository/models"
)

// Интерфейс AccountRepository описывает методы взаимодействия уровня service с уровнем repository
type AccountRepository interface {
	// UserByID получает данные пользователя по указанному ID
	UserByID(ctx context.Context, userID string) (*repModels.User, error)

	// Subscriptions получает подписки пользователя по указанному ID
	SubscriptionsByID(ctx context.Context, userID string) ([]repModels.Subscription, error)

	// AvatarPathByID получает путь до аватара пользователя по указанному ID
	AvatarPathByID(ctx context.Context, userID string) (string, error)

	// UpdateUsername обновляет имя пользователя
	UpdateUsername(ctx context.Context, userID string, username string) error

	// UpdatePassword обновляет пароль пользователя
	UpdatePassword(ctx context.Context, userID string, hashPassword string) error

	// GetPasswordHashByID получает хэш пароля пользователя по его ID
	GetPasswordHashByID(ctx context.Context, userID string) (string, error)

	// UpdateEmail обновляет почту пользователя
	UpdateEmail(ctx context.Context, userID string, email string) error

	// IsReader возвращает true, если пользователь является "reader"
	IsReader(ctx context.Context, userID string) (bool, error)

	// UpdateRoleToAuthor меняет роль пользователя на "author"
	UpdateRoleToAuthor(ctx context.Context, userID string) error

	// InitPage создаёт новую страницу автора после смены роли с "Reader" на "Author"
	InitPage(ctx context.Context, userID string) error

	// DeleteAvatar удаляет старый аватар пользователя при его обновлении
	DeleteAvatar(ctx context.Context, userID string) error

	// UpdateAvatar обновляет путь к аватару пользователя
	UpdateAvatar(ctx context.Context, userID string, file []byte, fileExtension string) error

	// GenerateID генерирует ID в формате UUIDv4
	GenerateID() string

	// NOTIFICATIONS

	// GetAllNotifications возвращает все уведомления c offset по (offset + limit)
	GetAllNotifications(ctx context.Context, userID string, offset, limit int) ([]*pkgModels.Notification, error)

	// GetNotReadNotifications возвращает непрочитанные  уведомления c offset по (offset + limit)
	GetNotReadNotifications(ctx context.Context, userID string, offset, limit int) ([]*pkgModels.Notification, error)

	// GetNewNotificationsByTime возвращает последние уведомления за указанное время
	GetNewNotificationsByTime(ctx context.Context, userID string, time int) ([]*pkgModels.Notification, error)

	// ChangeNotificationStatus - изменяет статус уведомления на "прочитано"
	ChangeNotificationStatus(ctx context.Context, userID, notificationID string) error
}
