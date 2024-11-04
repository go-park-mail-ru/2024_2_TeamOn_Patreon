package interfaces

import (
	"context"
	"mime/multipart"

	cModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/controller/models"
)

// Интерфейс AuthorService необходим для взаимодействия уровня controller с уровнем service
type AuthorService interface {
	// GetAuthorDataByID - получение данных автора по authorID
	GetAuthorDataByID(ctx context.Context, authorID string) (cModels.Author, error)

	// GetBackgroundByID - получение фона страницы автора по authorID
	GetBackgroundByID(ctx context.Context, authorID string) ([]byte, error)

	// PostUpdateInfo - обновление информации о себе
	PostUpdateInfo(ctx context.Context, authorID string, info string) error

	// GetAuthorPayments - получение выплат автора
	GetAuthorPayments(ctx context.Context, authorID string) (cModels.Payments, error)

	// PostUpdateBackground - изменение фона страницы автора
	PostUpdateBackground(ctx context.Context, authorID string, avatar multipart.File, fileName string) error
}
