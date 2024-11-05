package main

import (
	"context"
	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/cmd/microservices/auth/api"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/config"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/repository/postgresql"
	behavior "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	middlewares "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/repository/postgres"
	"net/http"
)

func main() {
	op := "cmd.microservices.auth.main"

	// logger
	logger.New()

	// config
	config.InitEnv("config/.env.default", "config/auth/.env.default")

	// repository
	db := postgres.InitPostgresDB(context.Background())
	defer db.Close()

	rep := postgresql.NewAuthRepository(db)

	// service
	beh := behavior.New(rep)

	// routers
	router := api.NewRouter(beh)

	// регистрируем middlewares
	router.Use(middlewares.Logging) // 1
	// router.Use(middlewares.HandlerAuth) // 2 только для ручек, где требуется аутентификация

	// run end-to-end
	port := config.GetEnv("SERVICE_PORT", "8081")
	logger.StandardInfoF(op, "Starting server at: %v", port)
	http.ListenAndServe(":"+port, router)
}
