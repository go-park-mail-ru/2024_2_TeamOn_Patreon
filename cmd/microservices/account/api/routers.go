package api

import (
	"context"
	"strings"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"net/http"

	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/controller"
	interfaces "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/controller/interfaces"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"

	metrics "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares/metrics"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string           // Имя ручки
	Method      string           // GET - POST
	Pattern     string           // Путь получения
	HandlerFunc http.HandlerFunc // Функция обработки ручки
}

type Routes []Route

func NewRouter(service interfaces.AccountService, monster *middlewares.Monster) *mux.Router {
	mainRouter := mux.NewRouter().StrictSlash(true)

	authRouter := mainRouter.PathPrefix("/").Subrouter()
	router := mainRouter.PathPrefix("/").Subrouter()

	handleAuth(authRouter, service)
	handleOther(router, service)

	authRouter.Use(monster.HandlerAuth)
	router.Use(monster.AuthMiddleware)

	mainRouter.Use(middlewares.CsrfMiddleware)
	mainRouter.Use(middlewares.Logging)
	mainRouter.Use(middlewares.AddRequestID)

	// Метрики
	metrics.NewMetrics(prometheus.DefaultRegisterer)
	mainRouter.Use(middlewares.MetricsMiddleware)

	return mainRouter
}

func handleAuth(router *mux.Router, service interfaces.AccountService) *mux.Router {
	op := "account.api.routers.NewRouterWithAuth"

	handler := api.New(service)

	var routes = Routes{
		Route{
			"GetAccount",
			"GET",
			"/account",
			handler.GetAccount,
		},
		Route{
			"PostAccountUpdate",
			"POST",
			"/account/update",
			handler.PostAccountUpdate,
		},
		Route{
			"PostAccountUpdateAvatar",
			"POST",
			"/account/update/avatar",
			handler.PostAccountUpdateAvatar,
		},
		Route{
			"PostAccountUpdateRole",
			"POST",
			"/account/update/role",
			handler.PostAccountUpdateRole,
		},
		Route{
			"GetNewNotifications",
			"GET",
			"/notification/new",
			handler.GetNewNotifications,
		},
		Route{
			"GetNotifications",
			"GET",
			"/notification",
			handler.GetNotifications,
		},
		Route{
			"PostNotificationStatusUpdate",
			"POST",
			"/notification/status/update",
			handler.PostNotificationStatusUpdate,
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

func handleOther(router *mux.Router, service interfaces.AccountService) *mux.Router {
	op := "account.api.routers.NewRouterOther"

	handler := api.New(service)

	var routes = Routes{
		Route{
			"GetAccountAvatar",
			"GET",
			"/account/{userID}/avatar",
			handler.GetAccountAvatar,
		},
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

	return router
}
