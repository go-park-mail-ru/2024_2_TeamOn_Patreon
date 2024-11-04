package controller

import (
	"net/http"

	global "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
)

// PostAuthorUpdateBackground - ручка изменения фона страницы автора
func (handler *Handler) PostAuthorUpdateBackground(w http.ResponseWriter, r *http.Request) {
	op := "internal.account.controller.PostAuthorUpdateBackground"

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Парсинг данных из multipart/form-data
	if err := r.ParseMultipartForm(5 << 20); err != nil { // 5 MB limit
		logger.StandardWarnF(op, "error parsing multipart form {%v}", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Проверяем, есть ли данные в форме
	if r.MultipartForm == nil || len(r.MultipartForm.File) == 0 {
		logger.StandardWarnF(op, "no files uploaded")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Извлекаем userData из контекста
	userData, ok := r.Context().Value(global.UserKey).(bModels.User)
	if !ok {
		logger.StandardResponse("userData not found in context", http.StatusUnauthorized, r.Host, op)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Валидация userID на соответствие стандарту UUIDv4
	if ok := utils.IsValidUUIDv4(string(userData.UserID)); !ok {
		logger.StandardResponse("invalid userID format", http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Получаем файл из формы
	fileBackground, fileHeader, err := r.FormFile("file")
	if err != nil {
		logger.StandardResponse("error retrieving file. key must be 'file'", http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer fileBackground.Close()

	// Обращение к service
	if err := handler.serv.PostUpdateBackground(r.Context(), string(userData.UserID), fileBackground, fileHeader.Filename); err != nil {
		logger.StandardWarnF(op, "update data error {%v}", err)
		// Status 500
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Status 200
	w.WriteHeader(http.StatusOK)
}
