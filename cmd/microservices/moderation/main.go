package main

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/cmd/microservices/moderation/api"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/config"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/moderation/repository/postgresql"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/moderation/service"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/repository/postgres"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"net/http"
)

func main() {
	op := "cmd.microservices.moderation.main"

	// logger
	logger.New()

	// pkg
	config.InitEnv("config/.env.default", "config/moderation/.env.default")

	// repository
	db := postgres.InitPostgresDB(context.Background())
	defer db.Close()

	rep := postgresql.NewModerationRepository(db)

	// service
	beh := service.New(rep)

	monster := middlewares.NewMonster()
	defer monster.Close()

	// handlers
	router := api.NewRouter(beh, monster)

	// run end-to-end
	port := config.GetEnv("SERVICE_PORT", "8087")
	logger.StandardInfoF(context.Background(), op, "Starting server at: %v", port)
	http.ListenAndServe(":"+port, router)
}
