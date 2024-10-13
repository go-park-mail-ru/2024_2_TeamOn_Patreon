package main

import (
	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/cmd/microservices/auth/api"
	repositories "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/repository/repositories"
	behavior "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	middlewares "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares"
	"net/http"
)

func main() {
	op := "cmd.microservices.auth.main"

	// config

	// logger
	logger.New()

	// repository
	rep := repositories.New()

	// service
	beh := behavior.New(rep)

	// routers
	router := api.NewRouter(beh)

	// регистрируем middlewares
	router.Use(middlewares.Logging) // 1
	// router.Use(middlewares.HandlerAuth) // 2 только для ручек, где требуется аутентификация

	// run end-to-end
	logger.StandardInfo("Starting server at: 8081", op)
	http.ListenAndServe(":8081", router)
}
