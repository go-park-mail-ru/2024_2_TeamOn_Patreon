package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (ts *TestServer) TestFailGetProfile(t *testing.T) {
	// Создаем новый GET запрос для неавторизованного пользователя
	req := ts.MakeRequest(t, "GET", "/profile", nil)
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

}
