package main

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/cmd/microservices/custom_subscribe/api"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/config"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/repository/postgresql"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/service"
	//"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/repository/postgresql"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/repository/postgres"
	"net/http"
)

func main() {
	op := "cmd.microservices.custom_subscription.main"

	// logger
	logger.New()

	// pkg
	config.InitEnv("config/.env.default", "config/custom_subscription/.env.default")

	// repository
	db := postgres.InitPostgresDB(context.Background())
	defer db.Close()

	rep := postgresql.NewCustomSubscriptionRepository(db)

	// service
	beh := service.New(rep)

	// handlers
	router := api.NewRouter(beh)

	// run end-to-end
	port := config.GetEnv("SERVICE_PORT", "8085")
	logger.StandardInfoF(context.Background(), op, "Starting server at: %v", port)
	http.ListenAndServe(":"+port, router)
}
