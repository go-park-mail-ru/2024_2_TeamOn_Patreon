package controller

import (
	"net/http"

	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/models/mapper"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	cModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	utils2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/gorilla/mux"
)

const APIPostID = "postID"

// PostMediaGet в get параметрах под mediaID будет UUID контента
func (h *Handler) PostMediaGet(w http.ResponseWriter, r *http.Request) {
	op := "internal.content.controller.PostMediaGet"

	ctx := r.Context()

	vars := mux.Vars(r)         // Извлекаем параметры из запроса
	postID := vars[postIDParam] // Получаем значение параметра "postID"
	if postID == "" {
		err := global.ErrBadRequest

		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op+" not postID")
		utils2.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	user, ok := r.Context().Value(global.UserKey).(cModels.User)
	if !ok {
		// проставляем http.StatusUnauthorized 401
		logger.StandardResponse(ctx, "userData not found in context", http.StatusUnauthorized, r.Host, op)
	}
	userID := user.UserID

	medias, err := h.b.GetFile(ctx, string(userID), postID)
	if err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		utils2.SendModel(&tModels.ModelError{Message: err.Error()}, w, op, ctx)
		return
	}

	tMedias := mapper.MapCommonMediaSToControllerMedias(medias)

	response := models.MediaResponse{
		PostID:       postID,
		MediaContent: tMedias,
	}

	w.WriteHeader(http.StatusOK)
	utils.SendModel(response, w, op, ctx)
}
