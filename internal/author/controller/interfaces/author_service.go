package interfaces

import (
	"context"
	"mime/multipart"

	pkgModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/pkg/models"
	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/service/models"
)

// AuthorService необходим для взаимодействия уровня controller с уровнем service
type AuthorService interface {
	// GetAuthorDataByID - получение данных автора по authorID
	GetAuthorDataByID(ctx context.Context, authorID string) (sModels.Author, error)

	// GetUserIsSubscribe - получение статуса подписки на автора
	GetUserIsSubscribe(ctx context.Context, authorID, userID string) (bool, error)

	// GetAuthorSubscriptions - получение подписок автора по authorID
	GetAuthorSubscriptions(ctx context.Context, authorID string) ([]sModels.Subscription, error)

	// GetBackgroundByID - получение фона страницы автора по authorID
	GetBackgroundByID(ctx context.Context, authorID string) ([]byte, error)

	// PostUpdateInfo - обновление информации о себе
	PostUpdateInfo(ctx context.Context, authorID, info string) error

	// GetAuthorPayments - получение выплат автора
	GetAuthorPayments(ctx context.Context, authorID string) (int, error)

	// PostUpdateBackground - изменение фона страницы автора
	PostUpdateBackground(ctx context.Context, authorID string, avatar multipart.File, fileName string) error

	// PostTip - обновление информации о себе
	PostTip(ctx context.Context, userID, authorID string, cost int, message string) error

	// SUBSCRIBE

	// GetCostSubscription получает стоимость подписки по уровню и сроку подписки
	GetCostSubscription(ctx context.Context, monthCount int, authorID string, layer int) (string, error)

	// CreateSubscriptionRequest создаёт пользовательский запрос на подписку на автора
	CreateSubscriptionRequest(ctx context.Context, subReq sModels.SubscriptionRequest) error

	// CreateSubscribeRequest реализует запрос пользователя на подписку на автора
	RealizeSubscriptionRequest(ctx context.Context, subReqID string, status bool, description, userID string) error

	// STATISTIC

	// GetStatistic - возвращает структуру с массивом точек по времени и массивом точек по количеству выборки
	// statParam - параметр, по которому собирается статистика
	GetStatistic(ctx context.Context, userID, time, statParam string) (*pkgModels.Graphic, error)
}
