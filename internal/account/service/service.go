// Бизнес-логика сервиса Account

package service

import (
	"context"
	"fmt"

	cModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/controller/models"
	interfaces "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/service/interfaces"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
)

type Service struct {
	rep interfaces.AccountRepository
}

func New(repository interfaces.AccountRepository) *Service {
	return &Service{repository}
}

// GetAccDataByID - получение данных аккаунта по userID
func (s *Service) GetAccDataByID(ctx context.Context, userID string) (cModels.Account, error) {
	op := "internal.account.service.GetAccDataByID"

	// данные пользователя в формате service model
	user, err := s.rep.UserByID(ctx, userID)
	if err != nil {
		logger.StandardDebugF(op, "fail get user: {%v}", err)
		return cModels.Account{}, err
	}

	logger.StandardInfo(
		fmt.Sprintf("successful get user=%v with userID='%v'", user.Username, user.UserID),
		op)

	accountData := cModels.Account{
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		// Subscriptions:
	}
	return accountData, nil
}
