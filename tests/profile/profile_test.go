package tests

import (
	"net/http"
	"testing"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/api/utils"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/behavior/jwt"
	"github.com/stretchr/testify/assert"
)

func (ts *TestServer) TestFailGetProfile(t *testing.T) {
	// Создаем новый GET запрос для неавторизованного пользователя
	req := ts.MakeRequest(t, "GET", "/profile", nil, nil)
	assert.Equal(t, http.StatusUnauthorized, req.StatusCode, "Expected status code 401")

}

func (ts *TestServer) TestGetProfile(t *testing.T) {
	// Готовый JWT
	var tokenStr jwt.TokenString
	tokenStr = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjUsInVzZXJuYW1lIjoicG9saW5hIiwicm9sZSI6MSwiaXNzIjoiYXV0aC1hcHAiLCJleHAiOjE3Mjc2MTg0ODUsImlhdCI6MTcyNzUzMjA4NX0.bLK14lLHDen3qRQPe3zTMpw1oAJ6DUChPEb6l1g8VZs"

	// Сохранение токена в куки
	cookie := utils.CreateCookie(tokenStr)

	// Создаем новый GET запрос для неавторизованного пользователя
	req := ts.MakeRequest(t, "GET", "/profile", nil, []*http.Cookie{&cookie})
	defer req.Body.Close()
	assert.Equal(t, http.StatusUnauthorized, req.StatusCode, "Expected status code 401")

}

// Запуск всех куки тестов
func TestAuthCookieHandlers(t *testing.T) {
	// test server
	ts := SetupTestServer()
	defer ts.TearDown()

	// Run cookie tests
	ts.TestFailGetProfile(t)
	ts.TestGetProfile(t)

}
