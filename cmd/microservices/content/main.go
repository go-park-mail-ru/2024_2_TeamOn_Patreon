package main

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/cmd/microservices/content/api"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/repository/imagine"
	behavior "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/service"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares"
	_ "github.com/gorilla/mux"
	"net/http"
)

func main() {
	op := "cmd.microservices.content.main"

	// config

	// logger
	logger.New()

	// repository
	rep := imagine.New()

	// service
	beh := behavior.New(rep)

	// routers
	// Создаем основной маршрутизатор

	router := api.NewRouter(beh)

	// регистрируем middlewares
	router.Use(middlewares.Logging) // 1
	// auth middleware registered in api.New

	// run end-to-end
	logger.StandardInfo("Starting server at: 8081", op)
	http.ListenAndServe(":8081", router)
}
