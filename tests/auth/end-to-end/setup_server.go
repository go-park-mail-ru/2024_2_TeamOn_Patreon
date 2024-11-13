package end_to_end

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/cmd/microservices/auth/api"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/config"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/repository/postgresql"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/repository/postgres"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestServer содержит общий сервер
type TestServer struct {
	server *httptest.Server
}

// SetupTestServer инициализирует сервер и роутер
func SetupTestServer() *TestServer {
	// logger
	logger.New()

	// pkg
	config.InitEnv("pkg/.env.default", "pkg/auth/.env.default")

	// repository
	db := postgres.InitPostgresDB(context.Background())
	defer db.Close()

	rep := postgresql.NewAuthRepository(db)
	beh := service.New(rep)
	router := api.NewRouter(beh)
	//commonHandler := middlewares.CreateMiddlewareWithCommonRepository(rep, service.New)

	// Добавляет миддлваре
	router.Use(middlewares.Logging)
	//router.Use(commonHandler)

	// Создание тестового сервера
	ts := httptest.NewServer(router)
	return &TestServer{server: ts}
}

// TearDown останавливает тестовые сервер
func (ts *TestServer) TearDown() {
	ts.server.Close()
}

// MakeRequest создает запрос к тестовому серверу
func (ts *TestServer) MakeRequest(t *testing.T, method, path string, body interface{}) *http.Response {
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
