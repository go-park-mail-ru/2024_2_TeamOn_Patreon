package tests

import (
	"net/http"
	"testing"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/api/utils"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/behavior/jwt"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/buisness/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/global"
	"github.com/stretchr/testify/assert"
)

func (ts *TestServer) TestContextNotExist(t *testing.T) {
	// Создаем новый GET запрос для неавторизованного пользователя
	req := ts.MakeRequest(t, "GET", "/profile", nil, nil)
	defer req.Body.Close()
	assert.Equal(t, http.StatusUnauthorized, req.StatusCode, "Expected status code 401")
}

func (ts *TestServer) TestInvalidID(t *testing.T) {
	// Создание JWT
	user := bModels.User{
		UserID:   -12,
		Username: "Great Gatsby",
		Role:     1,
	}
	tokenStr, _ := jwt.CreateJWT(user, global.TTL)

	// Сохранение токена в куки
	cookie := utils.CreateCookie(tokenStr)

	// Создаем новый GET запрос для авторизованного пользователя
	req := ts.MakeRequest(t, "GET", "/profile", nil, []*http.Cookie{&cookie})
	defer req.Body.Close()
	assert.Equal(t, http.StatusBadRequest, req.StatusCode, "Expected status code 400")
}

func (ts *TestServer) TestGetProfile(t *testing.T) {
	// Создание JWT
	user := bModels.User{
		UserID:   125,
		Username: "Great Gatsby",
		Role:     1,
	}
	tokenStr, _ := jwt.CreateJWT(user, global.TTL)

	// Сохранение токена в куки
	cookie := utils.CreateCookie(tokenStr)

	// Создаем новый GET запрос для авторизованного пользователя
	req := ts.MakeRequest(t, "GET", "/profile", nil, []*http.Cookie{&cookie})
	defer req.Body.Close()
	assert.Equal(t, http.StatusOK, req.StatusCode, "Expected status code 200")

}

// Запуск всех куки тестов
func TestAuthCookieHandlers(t *testing.T) {
	// test server
	ts := SetupTestServer()
	defer ts.TearDown()

	// Run cookie tests
	ts.TestContextNotExist(t)
	ts.TestInvalidID(t)
	ts.TestGetProfile(t)

}
