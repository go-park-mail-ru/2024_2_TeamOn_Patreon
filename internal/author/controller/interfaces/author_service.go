package interfaces

import (
	"context"

	cModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/controller/models"
)

// Интерфейс AuthorService необходим для взаимодействия уровня controller с уровнем service
type AuthorService interface {
	// GetAuthorDataByID - получение данных автора по authorID
	GetAuthorDataByID(ctx context.Context, authorID string) (cModels.Author, error)
}
