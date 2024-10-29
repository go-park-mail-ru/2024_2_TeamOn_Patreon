package service

import (
	"context"
	"fmt"

	apiModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/controller/models"
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
func (s *Service) GetAccDataByID(ctx context.Context, userID string) (apiModels.Account, error) {
	op := "internal.account.service.GetAccDataByID"

	accountData := apiModels.Account{}
	// Достаём данные Account из DB по userID

	// --------- Здесь начинается репозиторий ---------
	// УБРАТЬ ОТСЮДА РЕПОЗИТОРИЙ, ПЕРЕНЕСТИ В service
	// rep := repository.Get()

	// Проверяем, что пользователь с userID существует
	isAccountExist, err := s.rep.AccountExist(ctx, userID)
	if !isAccountExist && err != nil {
		return accountData, err
	}

	// Если существует, то собираем его данные

	// Если такой записи нет, значит профиль новый, поэтому создаём новую запись в БД
	// Иначе возвращаем существующий профиль с запрашиваемым userID
	// if !isUserExist {
	// 	account, err = rep.SaveAccount(userID, userData.Username, userData.Role)
	// 	if err != nil {
	// 		// проставляем http.StatusInternalServerError
	// 		logger.StandardResponse(err.Error(), http.StatusInternalServerError, r.Host, op)
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		return
	// 	}
	// 	logger.StandardResponse(
	// 		fmt.Sprintf("create new account user=%v with userID='%v'", userData.Username, userData.UserID),
	// 		http.StatusOK, r.Host, op)
	// } else {
	// if isAccountExist {
	accountRepData, err := s.rep.FindByID(ctx, userID)
	// Если не удалось найти аккаунт
	if err != nil {
		return accountData, err
	}
	// }
	logger.StandardInfo(
		fmt.Sprintf("successful get user=%v with userID='%v'", accountRepData.Username, accountRepData.UserID),
		op)

	// --------- Здесь заканчивается репозиторий ---------
	accountData = apiModels.Account{
		Username: accountRepData.Username,
		Email:    accountRepData.Email,
		Role:     accountRepData.Role,
		// Subscriptions: accountRepData.Subscriptions,
	}
	return accountData, nil
}
