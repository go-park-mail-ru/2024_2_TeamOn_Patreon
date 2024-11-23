package main

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/cmd/microservices/csat/api"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/config"

	//"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/repository/postgresql"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
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
	//db := postgres.InitPostgresDB(context.Background())
	//defer db.Close()

	//rep := postgresql.NewContentRepository(db)

	// service
	//beh := behavior.New(rep)

	// handlers
	router := api.NewRouter(nil)

	// run end-to-end
	port := config.GetEnv("SERVICE_PORT", "8086")
	logger.StandardInfoF(context.Background(), op, "Starting server at: %v", port)
	http.ListenAndServe(":"+port, router)
}
