package end_to_end

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

// Тест регистрации
func (ts *TestServer) TestRegisterUser(t *testing.T) {
	// Тело запроса для регистрации
	reqBody := map[string]string{
		"username": "testuser",
		"password": "Test@1234",
	}

	// Создание POST request к /auth/register
	resp := ts.MakeRequest(t, "POST", "/auth/register", reqBody)
	defer resp.Body.Close()

	// Статус ок
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")

	// Проверка наличие куки
	cookies := resp.Cookies()
	assert.NotEmpty(t, cookies, "Expected JWT cookie, but no cookies received")
}

// Test входа по логину
func (ts *TestServer) TestLoginUser(t *testing.T) {
	// Тело запроса для логина
	reqBody := map[string]string{
		"username": "testuser",  // Такой же username как при регистрации
		"password": "Test@1234", // Такой жэ password
	}

	// Создание POST request к /auth/login
	resp := ts.MakeRequest(t, "POST", "/auth/login", reqBody)
	defer resp.Body.Close()

	// Статус ОК
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")

	// Проверка наличие куки
	cookies := resp.Cookies()
	assert.NotEmpty(t, cookies, "Expected JWT cookie, but no cookies received")
}

// Тест некорректной регистрации (username already exists)
func (ts *TestServer) TestRegisterUserConflict(t *testing.T) {
	// Тело запроса с таким же пользователем
	reqBody := map[string]string{
		"username": "testuser", // Тот же username
		"password": "Test@1234",
	}

	// Создание POST request к /auth/register
	resp := ts.MakeRequest(t, "POST", "/auth/register", reqBody)
	defer resp.Body.Close()

	//  400 Bad Request
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Expected status code 400")
}

// Тест на неверный пароль в login
func (ts *TestServer) TestLoginInvalidPassword(t *testing.T) {
	// Тело запроса с неверным паролем
	reqBody := map[string]string{
		"username": "testuser",   // Существующий пользователь
		"password": "WrongPass1", // Неверный пароль
	}

	// Создание POST request к /auth/login
	resp := ts.MakeRequest(t, "POST", "/auth/login", reqBody)
	defer resp.Body.Close()

	// Статус 400
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Expected status code 400")
}

// Запуск всех тестов
func TestAuthHandlers(t *testing.T) {
	// >> go test ./...

	// test end-to-end
	ts := SetupTestServer()
	defer ts.TearDown()

	// Run tests
	ts.TestRegisterUser(t)
	ts.TestLoginUser(t)
	ts.TestRegisterUserConflict(t)
	ts.TestLoginInvalidPassword(t)
}
