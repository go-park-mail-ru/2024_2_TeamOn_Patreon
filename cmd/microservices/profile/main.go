package main

import (
	slog "log/slog"
	"net/http"
	"os"

	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/dan-profile/cmd/microservices/profile/api"
)

func main() {
	// config

	// utils
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// storage

	// behavior

	// routers

	router := api.NewRouter()

	// run server
	slog.Info("starting server at: 8080")
	http.ListenAndServe(":8080", router)
}
