package interfaces

import (
	"context"
	"mime/multipart"

	repModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/repository/models"
)

// Интерфейс AuthorRepository описывает методы взаимодействия уровня service с уровнем repository
type AuthorRepository interface {
	// AuthorByID получает данные автора по указанному ID
	AuthorByID(ctx context.Context, authorID string) (*repModels.Author, error)

	// UserIsSubscribe - получение статуса подписки на автора
	UserIsSubscribe(ctx context.Context, authorID, userID string) (bool, error)

	// Subscriptions получает подписки автора по указанному ID
	SubscriptionsByID(ctx context.Context, authorID string) ([]repModels.Subscription, error)

	// UpdateInfo - обновление поля "О себе"
	UpdateInfo(ctx context.Context, authorID string, info string) error

	// Payments - получение суммы выплат автора за донаты и подписки
	Payments(ctx context.Context, authorID string) (int, error)

	// BackgroundPathByID получает путь до фона страницы автора
	BackgroundPathByID(ctx context.Context, authorID string) (string, error)

	// DeleteBackground удаляет старый фон страницы автора при его обновлении
	DeleteBackground(ctx context.Context, authorID string) error

	// UpdateBackground обновляет путь к фону страницы автора
	UpdateBackground(ctx context.Context, authorID string, background multipart.File, fileName string) error

	// NewTip сохраняет запись о пожертвовании
	NewTip(ctx context.Context, userID, authorID string, cost int, message string) error

	// CreateSubscribeRequest создаёт запрос на подписку
	CreateSubscribeRequest(ctx context.Context, subReq repModels.SubscriptionRequest) (string, error)

	// RealizeSubscribeRequest реализует запрос на подписку
	RealizeSubscribeRequest(ctx context.Context, subReqID string) (string, error)

	// GenerateID генерирует ID в формате UUIDv4
	GenerateID() string

	// GetUsername - получение имени пользователя по userID
	GetUsername(ctx context.Context, userID string) (string, error)

	// GetCustomSubscriptionInfo - получает основные данные о кастомной подписке
	GetCustomSubscriptionInfo(ctx context.Context, customSubID string) (string, string, error)

	// SendNotificationOfNewSubscriber - отправляет уведомление о новом подписчике
	SendNotificationOfNewSubscriber(ctx context.Context, message, userID, authorID string) error
}
