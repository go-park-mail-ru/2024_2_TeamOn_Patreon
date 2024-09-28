package api

import (
	"fmt"

	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/profile/api"
	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/profile/utils"

	"log/slog"
	"strings"

	// The "net/http" library has methods to implement HTTP clients and servers
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string           // Имя ручки
	Method      string           // GET - POST
	Pattern     string           // Путь получения
	HandlerFunc http.HandlerFunc // Функция обработки ручки
}

type Routes []Route

func NewRouter() *mux.Router {
	op := "profile.routers.NewRouter"

	// Declare a new router
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		slog.Info(fmt.Sprintf("Registered: %s %s | in %s", route.Method, route.Pattern, op))

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(utils.Logging(handler, route.Name)) // Замыкание для сигнала: "дёрнули ручку"
	}

	return router
}

var routes = Routes{
	Route{
		"ProfileGet",
		strings.ToUpper("Get"),
		"/profile",
		api.ProfileGet,
	},
}
