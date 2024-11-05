package main

import (
	"net/http"

	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/cmd/microservices/author/api"
	postgres "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/init/postgres_db"
	repositories "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/repository"
	service "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/service"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	middlewares "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares"
)

func main() {
	op := "cmd.microservices.author.main"

	// logger
	logger.New()

	// connect to DB
	pool := postgres.InitPostgres()

	defer pool.Close()

	// repository
	rep := repositories.New(pool)

	// service
	serv := service.New(rep)

	// routers
	router := api.NewRouter(serv)

	// регистрируем middlewares
	router.Use(middlewares.Logging)     // 1
	router.Use(middlewares.HandlerAuth) // 2 для ручек, где требуется аутентификация
	router.Use(middlewares.CsrfMiddleware)

	// run server
	logger.StandardInfo("Starting server at: 8083", op)
	http.ListenAndServe(":8083", router)
}
