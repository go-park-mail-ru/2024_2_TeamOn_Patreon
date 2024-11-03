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
	ctx := context.Background()
	db, err := postgres.InitPostgresDB(ctx)
	if err != nil {
		logger.StandardError("panic: "+err.Error(), "main")
		logger.StandardDebugF(op, "PANIC = %v", err)
		return
	}
	rep := postgresql.NewAuthRepository(db)

	// service
	beh := behavior.New(rep)

	// routers
	router := api.NewRouter(beh)

	// регистрируем middlewares
	router.Use(middlewares.Logging) // 1
	// router.Use(middlewares.HandlerAuth) // 2 только для ручек, где требуется аутентификация

	// run end-to-end
	logger.StandardInfo("Starting server at: 8082", op)
	http.ListenAndServe(":8082", router)
}
