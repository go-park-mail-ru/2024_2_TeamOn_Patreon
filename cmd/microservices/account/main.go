package main

import (
	"net/http"

	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/cmd/microservices/account/api"
	postgres "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/postgres_db"
	repositories "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/repository"
	service "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/service"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	middlewares "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares"
)

func main() {
	op := "cmd.microservices.profile.main"

	// logger
	logger.New()

	// connect to DB
	connect := postgres.InitPostgres()

	// repository
	rep := repositories.New(connect)

	// service
	serv := service.New(rep)

	// routers
	router := api.NewRouter(serv)

	// регистрируем middlewares
	router.Use(middlewares.Logging)     // 1
	router.Use(middlewares.HandlerAuth) // 2 для ручек, где требуется аутентификация

	// run server
	logger.StandardInfo("Starting server at: 8082", op)
	http.ListenAndServe(":8082", router)
}
