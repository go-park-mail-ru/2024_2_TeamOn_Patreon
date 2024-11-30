package api

import (
	"context"
	"strings"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"net/http"

	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/controller"
	interfaces "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/controller/interfaces"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string           // Имя ручки
	Method      string           // GET - POST
	Pattern     string           // Путь получения
	HandlerFunc http.HandlerFunc // Функция обработки ручки
}

type Routes []Route

func NewRouter(behavior interfaces.AuthorService, monster *middlewares.Monster) *mux.Router {
	mainRouter := mux.NewRouter().StrictSlash(true)

	authRouter := mainRouter.PathPrefix("/").Subrouter()
	router := mainRouter.PathPrefix("/").Subrouter()

	handleAuth(authRouter, behavior)
	handleOther(router, behavior)

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

func handleAuth(router *mux.Router, service interfaces.AuthorService) *mux.Router {
	op := "author.api.routers.NewRouterWithAuth"

	handler := api.New(service)

	var routes = Routes{
		Route{
			"GetAuthorPayments",
			"GET",
			"/author/payments",
			handler.GetAuthorPayments,
		},
		Route{
			"PostAuthorUpdateInfo",
			"POST",
			"/author/update/info",
			handler.PostAuthorUpdateInfo,
		},
		Route{
			"PostAuthorUpdateBackground",
			"POST",
			"/author/update/background",
			handler.PostAuthorUpdateBackground,
		},
		Route{
			"PostAuthorTip",
			"POST",
			"/author/{authorID}/tip",
			handler.PostAuthorTip,
		},
		Route{
			"SubscriptionRequest",
			"POST",
			"/subscription/request",
			handler.PostSubscriptionRequest,
		},
		Route{
			"SubscriptionRealize",
			"POST",
			"/subscription/realize",
			handler.PostSubscriptionRealize,
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

func handleOther(router *mux.Router, service interfaces.AuthorService) *mux.Router {
	op := "author.api.routers.NewRouterOther"

	handler := api.New(service)

	var routes = Routes{
		Route{
			"GetAuthorBackground",
			"GET",
			"/author/{authorID}/background",
			handler.GetAuthorBackground,
		},
		Route{
			"GetAuthor",
			"GET",
			"/author/{authorID}",
			handler.GetAuthor,
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
