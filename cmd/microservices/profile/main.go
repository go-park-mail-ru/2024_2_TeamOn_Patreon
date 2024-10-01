package main

import (
	"net/http"

	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/cmd/microservices/profile/api"
	behavior "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/behavior"
	repositories "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/repository/repositories"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/logger"
	middlewares "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/middlewares"
)

func main() {
	op := "cmd.microservices.profile.main"

	// config

	// logger
	logger.New()

	// repository
	rep := repositories.New()

	// routers
	router := api.NewRouter()

	// регистрируем middlewares
	commonHandler := middlewares.CreateMiddlewareWithCommonRepository(rep, behavior.New)
	router.Use(middlewares.Logging)     // 1
	router.Use(middlewares.HandlerAuth) // 2 только для ручек, где требуется аутентификация
	router.Use(commonHandler)           // 3

	// run server
	logger.StandardInfo("Starting server at: 8082", op)
	http.ListenAndServe(":8082", router)
}
