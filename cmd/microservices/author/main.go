package main

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/config"

	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/cmd/microservices/author/api"
	postgres "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/init/postgres_db"
	repositories "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/repository"
	service "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/service"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
)

func main() {
	op := "cmd.microservices.author.main"

	// logger
	logger.New()

	config.InitEnv("config/.env.default", "config/author/.env.default")

	// connect to DB
	pool := postgres.InitPostgres(context.Background())

	defer pool.Close()

	// repository
	rep := repositories.New(pool)

	// service
	serv := service.New(rep)

	monster := middlewares.NewMonster()
	defer monster.Close()

	// routers
	router := api.NewRouter(serv, monster)

	// run server
	port := config.GetEnv("SERVICE_PORT", "8083")
	logger.StandardInfoF(context.Background(), op, "Starting server at: %v", port)
	http.ListenAndServe(":"+port, router)
}
