package interfaces

import (
	"context"

	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/service/models"
)

// Интерфейс AuthorRepository описывает методы взаимодействия уровня service с уровнем repository
type AuthorRepository interface {
	// AuthorByID получает данные автора по указанному ID
	AuthorByID(ctx context.Context, authorID string) (*sModels.Author, error)

	// GenerateID генерирует ID в формате UUIDv4
	// GenerateID() string
}
