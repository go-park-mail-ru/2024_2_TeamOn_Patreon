package tests

import (
	"net/http"
	"testing"

	repositories "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/repository/repositories"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller/utils"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service/jwt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/stretchr/testify/assert"
)

func (ts *TestServer) TestContextNotExist(t *testing.T) {
	// Создаем новый GET запрос для неавторизованного пользователя
	req := ts.MakeRequest(t, "GET", "/account", nil, nil)
	defer req.Body.Close()
	assert.Equal(t, http.StatusUnauthorized, req.StatusCode, "Expected status code 401")
}

func (ts *TestServer) TestInvalidID(t *testing.T) {
	// Создание JWT
	user := bModels.User{
		UserID:   "-12",
		Username: "Great Gatsby",
		Role:     1,
	}
	tokenStr, _ := jwt.CreateJWT(user, global.TTL)

	// Сохранение токена в куки
	cookie := utils.CreateCookie(tokenStr)

	// Создаем новый GET запрос для авторизованного пользователя
	req := ts.MakeRequest(t, "GET", "/account", nil, []*http.Cookie{&cookie})
	defer req.Body.Close()
	assert.Equal(t, http.StatusBadRequest, req.StatusCode, "Expected status code 400")
}

func (ts *TestServer) TestGetAccountFromReg(t *testing.T) {
	// Создание JWT
	user := bModels.User{
		UserID:   "9dc785df-17e3-43b9-9475-cf26ec4ac08b",
		Username: "maxround",
		Role:     "4bfa776c-3048-4291-8479-2a31a07f074f",
	}
	tokenStr, _ := jwt.CreateJWT(user, global.TTL)

	// Сохранение токена в куки
	cookie := utils.CreateCookie(tokenStr)

	// Создаем новый GET запрос для авторизованного пользователя
	req := ts.MakeRequest(t, "GET", "/account", nil, []*http.Cookie{&cookie})
	defer req.Body.Close()
	assert.Equal(t, http.StatusOK, req.StatusCode, "Expected status code 200")
}

func (ts *TestServer) TestGetAccountFromAuth(t *testing.T) {
	// Тест на GET существующего пользователя
	rep := repositories.New()
	rep.SaveAccount("9dc785df-17e3-43b9-9475-cf26ec4ac08b", "Great Gatsby", 1)

	// Создание JWT
	user := bModels.User{
		UserID:   "9dc785df-17e3-43b9-9475-cf26ec4ac08b",
		Username: "Great Gatsby",
		Role:     1,
	}
	tokenStr, _ := jwt.CreateJWT(user, global.TTL)

	// Сохранение токена в куки
	cookie := utils.CreateCookie(tokenStr)

	// Создаем новый GET запрос для авторизованного пользователя
	req := ts.MakeRequest(t, "GET", "/account", nil, []*http.Cookie{&cookie})
	defer req.Body.Close()
	assert.Equal(t, http.StatusOK, req.StatusCode, "Expected status code 200")
}

// Запуск всех GET профиль тестов
func TestAccount(t *testing.T) {
	// test server
	ts := SetupTestServer()
	defer ts.TearDown()

	t.Run("Context does not exist", func(t *testing.T) {
		ts.TestContextNotExist(t)
	})

	t.Run("Invalid ID", func(t *testing.T) {
		ts.TestInvalidID(t)
	})

	t.Run("Get account for exist (auth) user", func(t *testing.T) {
		ts.TestGetAccountFromReg(t)
	})

	t.Run("Get account for new (reg) user", func(t *testing.T) {
		ts.TestGetAccountFromAuth(t)
	})

}
