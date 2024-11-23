package main

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/cmd/microservices/csat/api"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/config"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/repository"
	behavior "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/service"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/repository/postgres"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/repository"
	behavior "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/service"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/repository/postgres"

	//"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/repository/postgres"
	"net/http"
)

func main() {
	op := "cmd.microservices.content.main"

	// logger
	logger.New()

	// pkg
	config.InitEnv("config/.env.default", "config/csat/.env.default")

	// repository
	db := postgres.InitPostgresDB(context.Background())
	defer db.Close()

	rep := repository.NewCSATRepository(db)

	// service
	beh := behavior.New(rep)

	// handlers
	router := api.NewRouter(beh)

	// run end-to-end
	port := config.GetEnv("SERVICE_PORT", "8086")
	logger.StandardInfoF(context.Background(), op, "Starting server at: %v", port)
	http.ListenAndServe(":"+port, router)
}
