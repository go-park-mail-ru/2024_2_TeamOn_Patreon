// This is the name of our package
// Everything with this package name can see everything
// else inside the same package, regardless of the file they are in
package api

import (
	"fmt"
	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/api"
	bInterfaces "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/behavior/interfaces"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/logger"
	"strings"

	// The "net/http" library has methods to implement HTTP clients and servers
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(behavior bInterfaces.AuthBehavior) *mux.Router {
	op := "auth.routers.NewRouter"

	handler := api.New(behavior)

	var routes = Routes{
		Route{
			"LoginPost",
			strings.ToUpper("Post"),
			"/auth/login",
			handler.LoginPost,
		},

		Route{
			"AuthRegisterPost",
			strings.ToUpper("Post"),
			"/auth/register",
			handler.AuthRegisterPost,
		},
	}

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
