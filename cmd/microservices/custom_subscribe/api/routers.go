package api

import (
	"context"
	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/controller"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/controller/interfaces"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares"
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

func NewRouter(behavior interfaces.CustomSubscriptionService) *mux.Router {
	mainRouter := mux.NewRouter().StrictSlash(true)

	authRouter := mainRouter.PathPrefix("/").Subrouter()
	router := mainRouter.PathPrefix("/").Subrouter()

	handleAuth(authRouter, behavior)
	handleOther(router, behavior)

	authRouter.Use(middlewares.HandlerAuth)
	router.Use(middlewares.AuthMiddleware)

	// регистрируем middlewares
	mainRouter.Use(middlewares.CsrfMiddleware)
	mainRouter.Use(middlewares.Logging)
	mainRouter.Use(middlewares.AddRequestID)

	return mainRouter
}

// handleAuth регистрирует ручки, где нужна аутентификация
func handleAuth(router *mux.Router, behavior interfaces.CustomSubscriptionService) *mux.Router {
	op := "custom_subscription.api.routers.handleAuth"

	handler := api.New(behavior)

	var routes = Routes{
		Route{
			// возвращает кастомные подписки автора
			"SubscriptionCustomPost",
			http.MethodPost,
			"/subscription/custom",
			handler.SubscriptionCustomPost,
		},
		Route{
			// возвращает уровни подписок автора, на которых у автора нет кастомных подписок
			"SubscriptionLayersGet",
			http.MethodGet,
			"/subscription/layers",
			handler.SubscriptionLayersGet,
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
func handleOther(router *mux.Router, behavior interfaces.CustomSubscriptionService) *mux.Router {
	op := "custom_subscription.api.routers.handleOther"

	handler := api.New(behavior)

	var routes = Routes{
		Route{
			// возвращает кастомные подписки автора
			"SubscriptionAuthorIDCustomGet",
			http.MethodGet,
			"/subscription/{" + api.PathAuthorID + "}/custom",
			handler.SubscriptionAuthorIDCustomGet,
		},
		Route{
			// возвращает список авторов по запросу поиска
			"SearchAuthor",
			http.MethodGet,
			"/search/{" + api.PathAuthorName + "}",
			handler.SearchAuthorNameGet,
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
