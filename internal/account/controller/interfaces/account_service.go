package interfaces

import (
	"context"
	"mime/multipart"

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

	// PostAccUpdateByID - изменение данных аккаунта по userID
	PostAccUpdateByID(ctx context.Context, userID, username, password, email string) error

	// PostAccountUpdateAvatar - изменение аватарки аккаунта по userID
	PostUpdateAvatar(ctx context.Context, userID string, avatar multipart.File, fileName string) error

	// PostUpdateRole - изменение аватарки аккаунта по userID
	PostUpdateRole(ctx context.Context, userID string) error
}
