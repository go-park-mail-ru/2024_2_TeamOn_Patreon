package interfaces

import (
	"context"

	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/service/models"
)

// Интерфейс AuthorRepository описывает методы взаимодействия уровня service с уровнем repository
type AuthorRepository interface {
	// AuthorByID получает данные автора по указанному ID
	AuthorByID(ctx context.Context, authorID string) (*sModels.Author, error)

	// UpdateInfo - обновление поля "О себе"
	UpdateInfo(ctx context.Context, authorID string, info string) error

	// Payments - получение суммы выплат автора за донаты и подписки
	Payments(ctx context.Context, authorID string) (int, error)

	// BackgroundPathByID получает путь до фона страницы автора
	BackgroundPathByID(ctx context.Context, authorID string) (string, error)

	// UpdateBackground обновляет путь к фону страницы автора
	UpdateBackground(ctx context.Context, userID string, backgroundPath string) error

	// GenerateID генерирует ID в формате UUIDv4
	GenerateID() string
}
