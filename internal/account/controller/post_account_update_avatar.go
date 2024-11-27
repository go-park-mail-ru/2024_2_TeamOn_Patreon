package controller

import (
	"net/http"

	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller/models"
	global "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/static"
	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
)

// PostAccountUpdateAvatar - ручка изменения аватарки пользователя
func (handler *Handler) PostAccountUpdateAvatar(w http.ResponseWriter, r *http.Request) {
	op := "internal.account.controller.PostAccountUpdateAvatar"

	ctx := r.Context()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

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

	// TODO: ограничение по весу

	// Получаем файл в формате multipart и его MIME-тип из запроса
	file, contentType, err := static.ExtractFileFromMultipart(r, "file")
	if err != nil {
		logger.StandardResponse(ctx, "error retrieving file. key must be 'file'", http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Получаем расширение файла на основе MIME-типа
	fileExtension, err := static.GetFileExtensionForPicture(contentType)
	if err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		// Status 415 - пользователь отправляет файл с недопустимым расширением
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// TODO: валидация на jpeg/jpg/png

	// Конвертируем multipart в []byte
	fileBytes, err := static.ConvertMultipartToBytes(file)
	if err != nil {
		logger.StandardWarnF(ctx, op, "error convert multipart to byte {%v}", err)
		// Status 500
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
	}

	// Обращение к service
	if err := handler.serv.PostUpdateAvatar(ctx, string(userData.UserID), fileBytes, fileExtension); err != nil {
		logger.StandardWarnF(ctx, op, "update data error {%v}", err)
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
	}

	// Status 200
	w.WriteHeader(http.StatusOK)
}
