package main

import (
	"auth/api"
	_ "auth/utils"
	"net/http"
	"os"

	"log/slog"
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
	slog.Info("starting server at: 8081")
	http.ListenAndServe(":8081", router)
}
