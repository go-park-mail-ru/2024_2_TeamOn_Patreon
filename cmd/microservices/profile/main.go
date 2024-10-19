package main

import (
	"net/http"

	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/cmd/microservices/profile/api"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	middlewares "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares"
)

func main() {
	op := "cmd.microservices.profile.main"

	// config

	// logger
	logger.New()

	// routers
	router := api.NewRouter()

	// регистрируем middlewares
	router.Use(middlewares.Logging)     // 1
	router.Use(middlewares.HandlerAuth) // 2 только для ручек, где требуется аутентификация

	// run server
	logger.StandardInfo("Starting server at: 8082", op)
	http.ListenAndServe(":8082", router)
}
