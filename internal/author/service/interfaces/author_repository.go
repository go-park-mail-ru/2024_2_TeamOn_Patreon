package interfaces

import (
	"context"
	"mime/multipart"

	pkgModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/pkg/models"
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

	// GenerateID генерирует ID в формате UUIDv4
	GenerateID() string

	// GetUsername - получение имени пользователя по userID
	GetUsername(ctx context.Context, userID string) (string, error)

	// GetCustomSubscriptionInfo - получает основные данные о кастомной подписке
	GetCustomSubscriptionInfo(ctx context.Context, customSubID string) (string, string, error)

	// SendNotification - отправляет уведомление
	SendNotification(ctx context.Context, message, userID, authorID string) error

	// GetAuthorPageID - получает ID страницы автора
	GetAuthorPageID(ctx context.Context, userID string) (string, error)

	// SUBSCRIPTION

	// GetCostCustomSub получает стоимость кастомной подписки по ID автора и уровню подписки
	GetCostCustomSub(ctx context.Context, authorID string, layer int) (int, error)

	// SaveSubscribeRequest сохраняет запрос на подписку
	SaveSubscribeRequest(ctx context.Context, subReq repModels.SubscriptionRequest) error

	// RealizeSubscribeRequest реализует запрос на подписку
	RealizeSubscribeRequest(ctx context.Context, subReqID string) (string, error)

	// TIP

	// NewTip сохраняет запись о пожертвовании [ УДАЛИТЬ КАК БУДЕТ ГОТОВО АПИ ОПЛАТЫ ]
	NewTip(ctx context.Context, userID, authorID string, cost int, message string) error

	// SaveTipRequest сохраняет запрос на пожертвование
	SaveTipRequest(ctx context.Context, tipReq repModels.TipRequest) error

	// STATISTIC

	// GetStatByDay - возвращает статистику за последний день по часам
	GetStatByDay(ctx context.Context, userID, statParam string) (*pkgModels.Graphic, error)

	// GetStatByMonth - возвращает статистику за последний месяц по дням
	GetStatByMonth(ctx context.Context, userID, statParam string) (*pkgModels.Graphic, error)

	// GetStatByYear - возвращает статистику за последний год по месяцам
	GetStatByYear(ctx context.Context, userID, statParam string) (*pkgModels.Graphic, error)
}
