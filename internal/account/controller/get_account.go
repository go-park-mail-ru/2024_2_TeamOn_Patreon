package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/controller/models"
	repModel "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/repository/models"
	repository "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/repository/repositories"
	global "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
)

// HandlerGetAccount - ручка получения данных профиля
func HandlerGetAccount(w http.ResponseWriter, r *http.Request) {
	op := "account.api.api_account"

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Извлекаем userData из контекста
	userData, ok := r.Context().Value(global.UserKey).(bModels.User)

	if !ok {
		// проставляем http.StatusUnauthorized 401
		logger.StandardResponse("userData not found in context", http.StatusUnauthorized, r.Host, op)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Валидация userID
	if !isValidUUIDv4(string(userData.UserID)) {
		// проставляем http.StatusBadRequest 400
		logger.StandardResponse("invalid userID format", http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Достаём данные Account из DB по userID
	// Проверяем, что пользователь с userID существует
	rep := repository.Get()
	isUserExist, err := rep.UserExist(string(userData.UserID))
	if err != nil {
		// проставляем http.StatusInternalServerError
		logger.StandardResponse(err.Error(), http.StatusInternalServerError, r.Host, op)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	account := &repModel.Account{}
	// Если такой записи нет, значит профиль новый, поэтому создаём новую запись в БД
	// Иначе возвращаем существующий профиль с запрашиваемым userID
	if !isUserExist {
		account, err = rep.SaveAccount(string(userData.UserID), userData.Username, userData.Role)
		if err != nil {
			// проставляем http.StatusInternalServerError
			logger.StandardResponse(err.Error(), http.StatusInternalServerError, r.Host, op)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		logger.StandardResponse(
			fmt.Sprintf("create new account user=%v with userID='%v'", userData.Username, userData.UserID),
			http.StatusOK, r.Host, op)
	} else {
		var err error
		account, err = rep.GetAccountByID(string(userData.UserID))
		// Если не удалось получить профиль
		if err != nil {
			// проставляем http.StatusInternalServerError
			logger.StandardResponse(err.Error(), http.StatusInternalServerError, r.Host, op)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	logger.StandardResponse(
		fmt.Sprintf("successful get user=%v with userID='%v'", userData.Username, userData.UserID),
		http.StatusOK, r.Host, op)
	accountData := models.Account{
		Username: account.Username,
		Email:    account.Email,
		Role:     account.Role,
		// Subscriptions: account.Subscriptions,
	}
	json.NewEncoder(w).Encode(accountData)
	w.WriteHeader(http.StatusOK)
}

func isValidUUIDv4(uuid string) bool {
	// Регулярное выражение для проверки формата UUID v4
	re := regexp.MustCompile(`^([0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12})$`)
	return re.MatchString(uuid)
}
