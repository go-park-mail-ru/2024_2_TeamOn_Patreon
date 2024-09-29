package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/buisness/models"
	global "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/global"
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
		slog.Info("err: userID not found in context")
		// TODO: Дописать отправку модели ошибки "Недопустимый ID пользователя" с err.msg
		return
	}

	// validate userID
	if userData.UserID <= 0 {
		slog.Info(fmt.Sprintf("err: incorrect format userID | in %v", op))
		// 	// TODO: Дописать отправку модели ошибки "Недопустимый ID пользователя" с err.msg
		return
	}

	// Достаём данные Profile из DB по userID
	// Проверяем, что пользователь с userID существует
	rep := repository.New()
	isUserExist, _ := rep.UserExists(userData.UserID)
	profile := &repModel.Profile{}
	// Если такой записи нет, значит профиль новый, поэтому создаём новую запись в БД
	// Иначе возвращаем профиль с запрашиваемым userID
	if !isUserExist {
		slog.Info(fmt.Sprintf("create new profile | in %v", op))
		profile, _ = rep.SaveProfile(userData.UserID, userData.Username, userData.Role)
	} else {
		slog.Info(fmt.Sprintf("profile found | in %v", op))
		var err error
		profile, err = rep.GetProfileByID(userData.UserID)
		if err != nil {
			slog.Info(fmt.Sprintf("error get profile | in %v", op))
			return
		}
		slog.Info(fmt.Sprintf("profile get | in %v", op))
	}

	// создаём объект Profile на основе полученных данных из  БД
	profileData := models.Profile{
		Username:      profile.Username,
		Email:         profile.Email,
		AvatarUrl:     profile.AvatarUrl,
		Role:          profile.Role,
		Followers:     profile.Followers,
		Subscriptions: profile.Subscriptions,
	}
	json.NewEncoder(w).Encode(profileData)
	w.WriteHeader(http.StatusOK)
}

func ProfilePaymentsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func ProfilePostsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// TODO: реализовать остальные ручки Profile согласно сваггеру
