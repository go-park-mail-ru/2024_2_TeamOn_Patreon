package server

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/behavior/jwt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/behavior/mapper"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/business/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/global"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func (ts *TestServer) TestUserRegisterWithCookie(t *testing.T) {
	// Тестовые данные для регистрации
	registerData := map[string]string{
		"username": "testuser",
		"password": "Test@1234",
	}

	// Отправляем запрос на регистрацию
	respRegister := ts.MakeRequest(t, "POST", "/auth/register", registerData)
	defer respRegister.Body.Close()

	// Проверяем успешный статус ответа
	assert.Equal(t, http.StatusOK, respRegister.StatusCode, "Регистрация неуспешна")

	// Проверка, что кука установлена
	cookies := respRegister.Cookies()
	require.Len(t, cookies, 1, "Должна быть установлена одна кука")

	cookie := cookies[0]
	assert.Equal(t, global.CookieJWT, cookie.Name, "Кука должна называться JWT")
	assert.NotEmpty(t, cookie.Value, "Кука должна содержать JWT-токен")
	assert.True(t, cookie.HttpOnly, "Кука должна быть HttpOnly")
	assert.Equal(t, "/", cookie.Path, "Путь куки должен быть /")
	assert.WithinDuration(t, time.Now().Add(global.TTL*time.Hour), cookie.Expires, time.Minute,
		"Время истечения куки должно быть корректным")

	// Создаем новый запрос для проверки токена и передаем туда куку
	req, err := http.NewRequest("GET", ts.server.URL+"/auth/check-token", nil)
	require.NoError(t, err, "Ошибка создания запроса для проверки токена")

	// Добавляем куки из ответа в новый запрос
	AddCookiesToRequest(req, respRegister.Cookies())

	// Тест распарсивания и валидации токена
	claims, err := jwt.ParseJWTFromCookie(req)
	require.NoError(t, err, "Ошибка парсинга токена из куки")

	// Проверяем корректность данных внутри токена
	assert.Equal(t, registerData["username"], claims.Username, "Имя пользователя должно совпадать с ожидаемым")
	assert.NotZero(t, claims.UserID, "ID пользователя должен быть задан")
	assert.Equal(t, bModels.Reader, claims.Role, "Роль пользователя должна быть 'user'")
}

func (ts *TestServer) TestUserLoginWithCookie(t *testing.T) {
	// Тестовые данные для логина
	loginData := map[string]string{
		"username": "testuser",
		"password": "Test@1234",
	}

	// Отправляем запрос на логин
	respLogin := ts.MakeRequest(t, "POST", "/auth/login", loginData)
	defer respLogin.Body.Close()

	// Проверяем успешный статус ответа
	assert.Equal(t, http.StatusOK, respLogin.StatusCode, "Логин неуспешен")

	// Проверяем, что кука с JWT установлена при логине
	cookiesLogin := respLogin.Cookies()
	require.Len(t, cookiesLogin, 1, "Должна быть установлена одна кука при логине")

	cookieLogin := cookiesLogin[0]
	assert.Equal(t, global.CookieJWT, cookieLogin.Name, "Кука должна называться JWT")
	assert.NotEmpty(t, cookieLogin.Value, "Кука должна содержать JWT-токен")
	assert.True(t, cookieLogin.HttpOnly, "Кука должна быть HttpOnly")
	assert.Equal(t, "/", cookieLogin.Path, "Путь куки должен быть /")
	assert.WithinDuration(t, time.Now().Add(global.TTL*time.Hour), cookieLogin.Expires, time.Minute,
		"Время истечения куки должно быть корректным")

	// Создаем новый запрос для проверки токена и передаем туда куку
	req, err := http.NewRequest("GET", ts.server.URL+"/auth/check-token", nil)
	require.NoError(t, err, "Ошибка создания запроса для проверки токена")

	// Добавляем куки из ответа в новый запрос
	AddCookiesToRequest(req, respLogin.Cookies())

	// Тест распарсивания и валидации токена после логина
	claimsLogin, err := jwt.ParseJWTFromCookie(req)
	require.NoError(t, err, "Ошибка парсинга токена из куки после логина")

	// Проверяем корректность данных внутри токена
	assert.Equal(t, loginData["username"], claimsLogin.Username, "Имя пользователя должно совпадать с ожидаемым")
	assert.NotZero(t, claimsLogin.UserID, "ID пользователя должен быть задан")
	assert.Equal(t, bModels.Reader, claimsLogin.Role, "Роль пользователя должна быть 'user'")

	// Маппинг токена в бизнес-модель пользователя
	user := mapper.MapTokenToUser(claimsLogin)
	assert.Equal(t, loginData["username"], user.Username, "Имя пользователя в бизнес-модели должно совпадать")
	assert.NotZero(t, user.UserID, "ID пользователя в бизнес-модели должен быть задан")
	assert.Equal(t, bModels.Reader, user.Role, "Роль пользователя в бизнес-модели должна быть 'user'")

}

// Запуск всех куки тестов
func TestAuthCookieHandlers(t *testing.T) {
	// test server
	ts := SetupTestServer()
	defer ts.TearDown()

	// Run cookie tests
	ts.TestUserRegisterWithCookie(t)
	ts.TestUserLoginWithCookie(t)

}
