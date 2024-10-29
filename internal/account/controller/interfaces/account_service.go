package interfaces

import (
	"context"

	apiModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/controller/models"
)

// Интерфейс AccountService необходим для взаимодействия уровня controller с уровнем service
type AccountService interface {
	// GetAccDataByID - получение данных аккаунта по userID
	GetAccDataByID(ctx context.Context, userID string) (apiModels.Account, error)
}
