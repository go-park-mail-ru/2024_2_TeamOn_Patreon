package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/buisness/models"
	global "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/logger"
	models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/profile/api/models"
	repModel "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/profile/repository/models"
	repository "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/profile/repository/repositories"
)

// ProfileGet - ручка получения данных профиля
func ProfileGet(w http.ResponseWriter, r *http.Request) {
	op := "profile.api.api_profile"
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
	if userData.UserID <= 0 {
		// проставляем http.StatusBadRequest 400
		logger.StandardResponse("invalid userID format", http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Достаём данные Profile из DB по userID
	// Проверяем, что пользователь с userID существует
	rep := repository.Get()
	isUserExist, err := rep.UserExist(userData.UserID)
	if err != nil {
		// проставляем http.StatusInternalServerError
		logger.StandardResponse(err.Error(), http.StatusInternalServerError, r.Host, op)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	profile := &repModel.Profile{}
	// Если такой записи нет, значит профиль новый, поэтому создаём новую запись в БД
	// Иначе возвращаем существующий профиль с запрашиваемым userID
	if !isUserExist {
		profile, err = rep.SaveProfile(userData.UserID, userData.Username, userData.Role)
		if err != nil {
			// проставляем http.StatusInternalServerError
			logger.StandardResponse(err.Error(), http.StatusInternalServerError, r.Host, op)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		logger.StandardResponse(
			fmt.Sprintf("create new profile user=%v with userID='%v'", userData.Username, userData.UserID),
			http.StatusOK, r.Host, op)
	} else {
		var err error
		profile, err = rep.GetProfileByID(userData.UserID)
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
	profileData := models.Profile{
		Username:      profile.Username,
		Email:         profile.Email,
		AvatarUrl:     profile.AvatarUrl,
		Status:        profile.Status,
		Role:          profile.Role,
		Followers:     profile.Followers,
		Subscriptions: profile.Subscriptions,
		PostsAmount:   profile.PostsAmount,
	}
	json.NewEncoder(w).Encode(profileData)
	w.WriteHeader(http.StatusOK)
}
