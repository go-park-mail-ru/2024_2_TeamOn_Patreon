package main

import (
	slog "log/slog"
	"net/http"
	"os"

	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/cmd/microservices/profile/api"
	behavior "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/behavior"
	repositories "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/repository/repositories"
	middlewares "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/middlewares"
)

func main() {
	// op := "cmd.microservices.profile.main"

	// config

	// logger [Переделать под оболочку logger]
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// repository
	rep := repositories.New()

	// routers
	router := api.NewRouter()

	// регистрируем middlewares
	commonHandler := middlewares.CreateMiddlewareWithCommonRepository(rep, behavior.New)
	router.Use(middlewares.Logging) // 1
	router.Use(commonHandler)       // 3

	// run server
	slog.Info("starting server at: 8080")
	http.ListenAndServe(":8080", router)
}
