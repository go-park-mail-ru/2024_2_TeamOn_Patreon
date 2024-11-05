package main

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/cmd/microservices/content/api"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/config"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/repository/postgresql"
	behavior "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/service"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/repository/postgres"
	_ "github.com/gorilla/mux"
	"net/http"
)

func main() {
	op := "cmd.microservices.content.main"

	// logger
	logger.New()

	// config
	config.InitEnv("config/.env.default", "config/content/.env.default")

	// repository
	db := postgres.InitPostgresDB(context.Background())
	defer db.Close()

	rep := postgresql.NewContentRepository(db)

	// service
	beh := behavior.New(rep)

	// routers
	// Создаем основной маршрутизатор

	router := api.NewRouter(beh)

	// регистрируем middlewares
	router.Use(middlewares.Logging) // 1
	// auth middleware registered in api.New

	// run end-to-end
	port := config.GetEnv("SERVICE_PORT", "8084")
	logger.StandardInfoF(op, "Starting server at: %v", port)
	http.ListenAndServe(":"+port, router)
}
