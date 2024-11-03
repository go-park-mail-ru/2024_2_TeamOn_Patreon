package interfaces

import (
	"context"
	"mime/multipart"

	cModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/controller/models"
)

// Интерфейс AccountService необходим для взаимодействия уровня controller с уровнем service
type AccountService interface {
	// GetAccDataByID - получение данных аккаунта по userID
	GetAccDataByID(ctx context.Context, userID string) (cModels.Account, error)

	// GetAvatarByID - получение аватарки пользователя по userID
	GetAvatarByID(ctx context.Context, userID string) ([]byte, error)

	// PostAccUpdateByID - изменение данных аккаунта по userID
	PostAccUpdateByID(ctx context.Context, userID string, username string, password string, email string) error

	// PostAccountUpdateAvatar - изменение аватарки аккаунта по userID
	PostUpdateAvatar(ctx context.Context, userID string, avatar multipart.File, fileName string) error

	// PostUpdateRole - изменение аватарки аккаунта по userID
	PostUpdateRole(ctx context.Context, userID string) error
}
