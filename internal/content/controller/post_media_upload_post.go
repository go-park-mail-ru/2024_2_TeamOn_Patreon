package controller

import (
	"net/http"

	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/static"
	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/gorilla/mux"
)

func (h *Handler) PostUploadMediaPost(w http.ResponseWriter, r *http.Request) {
	op := "content.controller.post_upload_content_post"

	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Достаем юзера
	user, ok := r.Context().Value(global.UserKey).(bModels.User)
	if !ok {
		// проставляем http.StatusUnauthorized 401
		logger.StandardResponse(ctx, "userData not found in context", http.StatusUnauthorized, r.Host, op)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Достаем postId
	vars := mux.Vars(r)
	postId, ok := vars[postIDParam]
	logger.StandardDebugF(ctx, op, "postId=%v", postId)
	if !ok {
		err := global.ErrBadRequest
		logger.StandardWarnF(ctx, op, "Received get post from path error={%v} post_id='%v'", err, postId)
		// проставляем http.StatusBadRequest
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	// Парсинг multipart/form
	err := r.ParseMultipartForm(32 << 20) // 32MB limit
	if err != nil {
		logger.StandardWarnF(ctx, op, "failed to parse multipart message {%v}", err)

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	logger.StandardDebugF(ctx, op, "want to get file keys from request body")

	// Получаем файлы с индексами
	files := r.MultipartForm.File
	if len(files) == 0 {
		logger.StandardResponse(ctx, "no files uploaded", http.StatusBadRequest, r.Host, op)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	logger.StandardDebugF(ctx, op, "got files: %v", files)
	// Извлекаем файлы из запроса
	// Здесь key == "file1", "file2", ..., "file[n]"
	for key, _ := range files {

		logger.StandardDebugF(ctx, op, "key %v in range files", key)
		// Получаем файл в формате multipart и его MIME-тип из запроса
		file, contentType, err := static.ExtractFileFromMultipart(r, key)
		if err != nil {
			logger.StandardWarnF(ctx, op, "error retrieving file {%v}", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Получаем расширение файла на основе MIME-типа
		fileExtension, err := static.GetFileExtension(contentType)
		if err != nil {
			logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
			// Status 415 - пользователь отправляет файл с недопустимым расширением
			w.WriteHeader(global.GetCodeError(err))
			utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
			return
		}

		// Конвертируем multipart в []byte
		fileBytes, err := static.ConvertMultipartToBytes(file)
		if err != nil {
			logger.StandardWarnF(ctx, op, "error convert multipart to byte {%v}", err)
			// Status 500
			w.WriteHeader(global.GetCodeError(err))
			utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		}

		// Обращение к service
		if err := h.b.UploadMedia(ctx, string(user.UserID), postId, fileBytes, fileExtension, key); err != nil {
			logger.StandardWarnF(ctx, op, "update data error {%v}", err)
			w.WriteHeader(global.GetCodeError(err))
			utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		}
	}

	w.WriteHeader(http.StatusOK)
}
