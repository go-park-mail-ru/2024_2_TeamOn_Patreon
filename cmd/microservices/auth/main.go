package main

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/config"
	"net/http"

	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/cmd/microservices/auth/api"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/repository/postgresql"
	behavior "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/repository/postgres"
)

func main() {
	op := "cmd.microservices.auth.main"

	// logger
	logger.New()

	// pkg
	config.InitEnv("config/.env.default", "config/auth/.env.default")

	// repository
	db := postgres.InitPostgresDB(context.Background())
	defer db.Close()

	rep := postgresql.NewAuthRepository(db)

	// service
	beh := behavior.New(rep)

	// routers
	router := api.NewRouter(beh)

	// run end-to-end
	port := config.GetEnv("SERVICE_PORT", "8081")
	logger.StandardInfoF(context.Background(), op, "Starting server at: %v", port)
	http.ListenAndServe(":"+port, router)
}
