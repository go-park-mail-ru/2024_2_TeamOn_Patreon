package main

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/cmd/microservices/content/api"
	behavior "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/service"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares"
	"net/http"
)

func main() {
	op := "cmd.microservices.content.main"

	// config

	// logger
	logger.New()

	// repository
	//rep := repositories.New()

	// service
	beh := behavior.New(nil)

	// routers
	router := api.NewRouter(beh)

	// регистрируем middlewares
	router.Use(middlewares.Logging) // 1
	//router.Use(middlewares.HandlerAuth) // 2 только для ручек, где требуется аутентификация

	// run end-to-end
	logger.StandardInfo("Starting server at: 8081", op)
	http.ListenAndServe(":8081", router)
}
