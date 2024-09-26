package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/dan-profile/internal/profile/api/models"
)

// ProfileGet - ручка получения данных профиля
func ProfileGet(w http.ResponseWriter, r *http.Request) {
	op := "profile.api.api_profile"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// *Откуда берем userID? Как параметр ручки, извлекаем из URL, парсим из кукисов?

	// parse userID
	userID := r.URL.Query().Get("user_id")
	// http://localhost:8080/profile?user_id=123

	// validate userID
	if userID == "" {
		return
	}
	userIDInt, err := strconv.Atoi(userID)
	if err != nil || userIDInt <= 0 {
		slog.Info(fmt.Sprintf("err: incorrect format userID | in %v", op))
		// TODO: Дописать отправку модели ошибки "Недопустимый ID пользователя" с err.msg
		return
	}

	// Find user in DB
	// For example:
	if userIDInt != 123 {
		slog.Info(fmt.Sprintf("err: profile does not exist | in %v", op))
		// TODO: Дописать отправку модели ошибки "Недопустимый ID пользователя" с err.msg
		return
	}

	slog.Info(fmt.Sprintf("profile found | in %v", op))

	// создаём объект Profile на основе полученных данных в БД
	profileData := models.Profile{
		Username:      "maxround",
		Email:         "blablaman@yandex.ru",
		AvatarUrl:     "https://example.com/avatar.jpg",
		Status:        models.AuthorStatus,
		Followers:     1140,
		Subscriptions: 2,
		Posts:         21,
	}
	json.NewEncoder(w).Encode(profileData)
	w.WriteHeader(http.StatusOK)
}

// TODO: реализовать остальные ручки Profile согласно сваггеру
