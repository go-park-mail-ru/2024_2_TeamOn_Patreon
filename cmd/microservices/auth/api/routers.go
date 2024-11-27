// This is the name of our package
// Everything with this package name can see everything
// else inside the same package, regardless of the file they are in
package api

import (
	"context"
	"fmt"

	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller"
	bInterfaces "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller/interafces"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	// The "net/http" library has methods to implement HTTP clients and servers
	"net/http"

	"github.com/gorilla/mux"
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
			"POST",
			"/auth/login",
			handler.LoginPost,
		},

		Route{
			"AuthRegisterPost",
			"POST",
			"/auth/register",
			handler.AuthRegisterPost,
		},

		Route{
			"LogoutPost",
			"POST",
			"/auth/logout",
			handler.LogoutPost,
		},
		Route{
			"GetCSRFToken",
			"GET",
			"/token-endpoint",
			middlewares.GetCSRFTokenHandler,
		},
		Route{
			"Metrics",
			"GET",
			"/metrics",
			promhttp.Handler().ServeHTTP,
		},
	}

	// Declare a new router
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		logger.StandardInfo(
			context.Background(),
			fmt.Sprintf("Registered: %s %s", route.Method, route.Pattern),
			op,
		)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	// регистрируем middlewares
	router.Use(middlewares.AddRequestID)
	router.Use(middlewares.Logging)
	router.Use(middlewares.Security)
	router.Use(middlewares.CsrfMiddleware)

	// Метрики
	metrics.NewMetrics(prometheus.DefaultRegisterer)
	router.Use(middlewares.MetricsMiddleware)

	return router
}
