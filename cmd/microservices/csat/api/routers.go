package api

import (
	"context"
	"net/http"
	"strings"

	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/controller"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/controller/interfaces"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares/metrics"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(behavior interfaces.CSATService) *mux.Router {
	mainRouter := mux.NewRouter().StrictSlash(true)

	authRouter := mainRouter.PathPrefix("/").Subrouter()
	router := mainRouter.PathPrefix("/").Subrouter()

	handleAuth(authRouter, behavior)
	handleOther(router)

	authRouter.Use(middlewares.HandlerAuth)

	// регистрируем middlewares
	mainRouter.Use(middlewares.CsrfMiddleware)
	mainRouter.Use(middlewares.Logging)
	mainRouter.Use(middlewares.AddRequestID)

	// Метрики
	metrics.NewMetrics(prometheus.DefaultRegisterer)
	mainRouter.Use(middlewares.MetricsMiddleware)

	return mainRouter
}

func handleAuth(router *mux.Router, behavior interfaces.CSATService) *mux.Router {
	op := "content.api.routers.handleAuth"

	handler := api.New(behavior)

	var routes = Routes{
		Route{
			"CsatCheckGet",
			strings.ToUpper("Get"),
			"/csat/check",
			handler.CsatCheckGet,
		},

		Route{
			"CsatQuestionGet",
			strings.ToUpper("Get"),
			"/csat/question",
			handler.CsatQuestionGet,
		},

		Route{
			"CsatResultQuestionIDPost",
			strings.ToUpper("Post"),
			"/csat/result/{questionID}",
			handler.CsatResultQuestionIDPost,
		},
		Route{
			"CsatTableGet",
			strings.ToUpper("Get"),
			"/csat/table",
			handler.CsatTableGet,
		},
	}

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		logger.StandardInfoF(context.Background(), op, "Registered: %s %s", route.Method, route.Pattern)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func handleOther(router *mux.Router) {
	op := "content.api.routers.NewRouterWithAuth"

	var routes = Routes{
		Route{
			"GetCSRFToken",
			strings.ToUpper("Get"),
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

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		logger.StandardInfoF(context.Background(), op, "Registered: %s %s", route.Method, route.Pattern)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
}
