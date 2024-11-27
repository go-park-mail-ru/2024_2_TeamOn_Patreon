package mock_interfaces

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service/jwt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func AddUserIDInReq(req *http.Request, userID string) *http.Request {
	// Новый контекст с добавленным значением
	ctx := context.WithValue(req.Context(), global.UserKey, bModels.User{UserID: bModels.UserID(userID)})

	// Новый запрос с обновленным контекстом
	req = req.WithContext(ctx)
	return req
}

func SetUserCookie(userIdStr string, req *http.Request) {
	user := bModels.User{UserID: bModels.UserID(userIdStr)}
	token, _ := jwt.CreateJWT(user, 10)

	req.AddCookie(&http.Cookie{Name: global.CookieJWT, Value: string(token)})
}

func GenerateUUID() string {
	return uuid.NewV4().String()
}

func AddValueToPath(req *http.Request, key, value string) *http.Request {
	req = mux.SetURLVars(req, map[string]string{key: value})
	return req
}
