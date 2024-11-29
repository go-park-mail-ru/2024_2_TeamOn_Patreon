package api

import (
	"context"
	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/moderation/controller"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/moderation/controller/interfaces"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares/metrics"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(behavior interfaces.ModerationService, monster *middlewares.Monster) *mux.Router {
	mainRouter := mux.NewRouter().StrictSlash(true)

	authRouter := mainRouter.PathPrefix("/").Subrouter()
	router := mainRouter.PathPrefix("/").Subrouter()

	handleAuth(authRouter, behavior)

	authRouter.Use(monster.HandlerAuth)
	router.Use(monster.AuthMiddleware)

	// регистрируем middlewares
	mainRouter.Use(middlewares.CsrfMiddleware)
	mainRouter.Use(middlewares.Logging)
	mainRouter.Use(middlewares.AddRequestID)

	// Метрики
	metrics.NewMetrics(prometheus.DefaultRegisterer)
	mainRouter.Use(middlewares.MetricsMiddleware)

	return mainRouter
}

func handleAuth(router *mux.Router, behavior interfaces.ModerationService) *mux.Router {
	op := "moderation.api.routers.handleAuth"
	handler := api.New(behavior)

	var routes = Routes{
		Route{
			"ModerationPostComplaintPost",
			http.MethodPost,
			"/moderation/post/complaint",
			handler.ModerationPostComplaintPost,
		},

		Route{
			"ModerationPostDecisionPost",
			http.MethodPost,
			"/moderation/post/decision",
			handler.ModerationPostDecisionPost,
		},

		Route{
			"ModerationPostGet",
			http.MethodGet,
			"/moderation/post",
			handler.ModerationPostGet,
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

// handleOther регистрирует ручки, где не нужна аутентификация
func handleOther(router *mux.Router, behavior interfaces.ModerationService) *mux.Router {
	op := "custom_subscription.api.routers.handleOther"

	var routes = Routes{
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

	return router
}
