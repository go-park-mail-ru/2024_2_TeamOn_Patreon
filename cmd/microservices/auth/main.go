package main

import (
	"context"
	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/cmd/microservices/auth/api"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/repository/postgresql"
	behavior "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	middlewares "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/repository/postgres"
	"net/http"
)

func main() {
	op := "cmd.microservices.auth.main"

	// config

	// logger
	logger.New()

	// repository
	db := postgres.InitPostgresDB(context.Background())

	rep := postgresql.NewAuthRepository(db)

	// service
	beh := behavior.New(rep)

	// routers
	router := api.NewRouter(beh)

	// регистрируем middlewares
	router.Use(middlewares.Logging) // 1
	// router.Use(middlewares.HandlerAuth) // 2 только для ручек, где требуется аутентификация

	// run end-to-end
	logger.StandardInfo("Starting server at: 8081", op)
	http.ListenAndServe(":8081", router)
}
