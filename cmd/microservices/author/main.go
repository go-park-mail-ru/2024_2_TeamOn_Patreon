package main

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/config"
	"net/http"

	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/cmd/microservices/author/api"
	postgres "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/init/postgres_db"
	repositories "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/repository"
	service "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/service"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	middlewares "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares"
)

func main() {
	op := "cmd.microservices.author.main"

	// logger
	logger.New()

	config.InitEnv("pkg/.env.default", "pkg/author/.env.default")

	// connect to DB
	pool := postgres.InitPostgres(context.Background())

	defer pool.Close()

	// repository
	rep := repositories.New(pool)

	// service
	serv := service.New(rep)

	// routers
	router := api.NewRouter(serv)

	// регистрируем middlewares
	router.Use(middlewares.AddRequestID)
	router.Use(middlewares.Logging)     // 1
	router.Use(middlewares.HandlerAuth) // 2 для ручек, где требуется аутентификация
	router.Use(middlewares.CsrfMiddleware)

	// run server
	port := config.GetEnv("SERVICE_PORT", "8083")
	logger.StandardInfoF(context.Background(), op, "Starting server at: %v", port)
	http.ListenAndServe(":"+port, router)
}
