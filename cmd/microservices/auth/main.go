package main

import (
	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/polina-auth/cmd/microservices/auth/api"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/polina-auth/internal/auth/behavior"
	repositories "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/polina-auth/internal/auth/repository/repositories"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/polina-auth/internal/common/logger"
	middlewares "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/polina-auth/internal/common/middlewares"
	"net/http"
)

func main() {
	op := "cmd.microservices.auth.main"

	// config

	// logger
	logger.New()

	// repository
	rep := repositories.New()

	// routers
	router := api.NewRouter()

	// регистрируем middlewares
	commonHandler := middlewares.CreateMiddlewareWithCommonRepository(rep, behavior.New)
	router.Use(middlewares.Logging) // 1
	// router.Use(middlewares.HandlerAuth) // 2 только для ручек, где требуется аутентификация
	router.Use(commonHandler) // 3

	// run server
	logger.StandardInfo("Starting server at: 8081", op)
	http.ListenAndServe(":8081", router)
}
