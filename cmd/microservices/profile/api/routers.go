package api

import (
	"fmt"

	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/logger"
	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/profile/api"

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
		logger.StandardInfo(
			fmt.Sprintf("Registered: %s %s", route.Method, route.Pattern),
			op,
		)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"ProfileGet",
		"GET",
		"/profile",
		api.ProfileGet,
	},
}
