package controller

import (
	"net/http"

	global "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
)

// PostAccountUpdateAvatar - ручка изменения аватарки пользователя
func (handler *Handler) PostAccountUpdateAvatar(w http.ResponseWriter, r *http.Request) {
	op := "internal.account.controller.PostAccountUpdateAvatar"

	ctx := r.Context()
	contentType := r.Header.Get("Content-Type")
	logger.StandardWarnF(ctx, op, "Content-Type: %s", contentType)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Парсинг данных из multipart/form-data
	if err := r.ParseMultipartForm(5 << 20); err != nil { // 5 MB limit
		logger.StandardWarnF(ctx, op, "error parsing multipart form {%v}", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Проверяем, есть ли данные в форме
	if r.MultipartForm == nil || len(r.MultipartForm.File) == 0 {
		logger.StandardWarnF(ctx, op, "no files uploaded")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Извлекаем userData из контекста
	userData, ok := r.Context().Value(global.UserKey).(bModels.User)
	if !ok {
		logger.StandardResponse(ctx, "userData not found in context", http.StatusUnauthorized, r.Host, op)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Валидация userID на соответствие стандарту UUIDv4
	if ok := utils.IsValidUUIDv4(string(userData.UserID)); !ok {
		logger.StandardResponse(ctx, "invalid userID format", http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Получаем файл из формы
	fileAvatar, fileHeader, err := r.FormFile("file")
	if err != nil {
		logger.StandardResponse(ctx, "error retrieving file. key must be 'file'", http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer fileAvatar.Close()

	// Обращение к service
	if err := handler.serv.PostUpdateAvatar(r.Context(), string(userData.UserID), fileAvatar, fileHeader.Filename); err != nil {
		logger.StandardWarnF(ctx, op, "update data error {%v}", err)
		// Status 500
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Status 200
	w.WriteHeader(http.StatusOK)
}
