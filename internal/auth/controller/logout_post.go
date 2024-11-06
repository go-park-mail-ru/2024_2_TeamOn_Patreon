package controller

import (
	"fmt"
	"net/http"

	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service/jwt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service/mapper"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
)

// LogoutPost - ручка разлогина
func (handler *Handler) LogoutPost(w http.ResponseWriter, r *http.Request) {
	op := "auth.controller.LogoutPost"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ctx := r.Context()

	// парсинг jwt токена
	tokenClaims, err := jwt.ParseJWTFromCookie(r)
	if err != nil || tokenClaims == nil {
		err = global.ErrUserNotAuthorized
		w.WriteHeader(global.GetCodeError(err))
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op)
		err = global.ErrUserNotAuthorized
		w.WriteHeader(global.GetCodeError(err))
		utils.SendStringModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op)
		return
	}

	// если все ок достаем юзер ид, юзернэйм и роль
	// мапим это все в структуру user для бизнес-логики
	user := mapper.MapTokenToUser(tokenClaims)

	err = handler.b.LogoutUser(ctx, user.UserID)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.SendStringModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op)
		return
	}
	// Сохранение токена в куки
	cookie := utils.CreateEmptyCookieJWT()

	// Устанавливаем токен в куку
	http.SetCookie(w, &cookie)

	logger.StandardResponse(
		ctx,
		fmt.Sprintf("Successful logout user=%v", user.Username),
		http.StatusOK, r.Host, op)

	w.WriteHeader(http.StatusOK)
}
