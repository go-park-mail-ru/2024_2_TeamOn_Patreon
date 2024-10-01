package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/cmd/microservices/profile/api"
	middlewares "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/middlewares"

	"github.com/stretchr/testify/assert"
)

// TestServer содержит общий сервер
type TestServer struct {
	server *httptest.Server
}

// SetupTestServer инициализирует сервер и роутер
func SetupTestServer() *TestServer {
	// routers
	router := api.NewRouter()

	// регистрируем middlewares
	router.Use(middlewares.Logging)
	router.Use(middlewares.HandlerAuth)

	// Создание тестового сервера
	ts := httptest.NewServer(router)
	return &TestServer{server: ts}
}

// TearDown останавливает тестовые сервер
func (ts *TestServer) TearDown() {
	ts.server.Close()
}

// MakeRequest создает запрос к тестовому серверу
func (ts *TestServer) MakeRequest(t *testing.T, method, path string, body interface{}, cookies []*http.Cookie) *http.Response {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		assert.NoError(t, err, "Error marshaling request")
	}

	req, err := http.NewRequest(method, ts.server.URL+path, bytes.NewBuffer(jsonBody))
	assert.NoError(t, err, "Error creating request")

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	AddCookiesToRequest(req, cookies)

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err, "Error sending request")
	return resp
}

// AddCookiesToRequest добавляет куки из response в запрос
func AddCookiesToRequest(req *http.Request, cookies []*http.Cookie) {
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
}
