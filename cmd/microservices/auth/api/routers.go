// This is the name of our package
// Everything with this package name can see everything
// else inside the same package, regardless of the file they are in
package api

import (
	"fmt"
	"log/slog"
	"strings"

	// The "net/http" library has methods to implement HTTP clients and servers
	"github.com/gorilla/mux"
	"net/http"

	"auth/utils"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	op := "auth.routers.NewRouter"

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
			Handler(utils.Logging(handler, route.Name))
	}

	return router
}

var routes = Routes{
	Route{
		"AuthLoginPost",
		strings.ToUpper("Post"),
		"/auth/login",
		AuthLoginPost,
	},

	Route{
		"AuthRegisterPost",
		strings.ToUpper("Post"),
		"/auth/register",
		AuthRegisterPost,
	},
}
