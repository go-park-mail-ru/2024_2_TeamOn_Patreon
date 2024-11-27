package api

import (
	"context"
	"net/http"
	"strings"

	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/interfaces"
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

func NewRouter(behavior interfaces.ContentBehavior) *mux.Router {
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

	// Метрики
	metrics.NewMetrics(prometheus.DefaultRegisterer)
	mainRouter.Use(middlewares.MetricsMiddleware)

	return mainRouter
}

func handleAuth(router *mux.Router, behavior interfaces.ContentBehavior) *mux.Router {
	op := "content.api.routers.NewRouterWithAuth"

	handler := api.New(behavior)

	var routes = Routes{
		Route{
			"PostLikePost",
			strings.ToUpper("Post"),
			"/post/like",
			handler.PostLikePost,
		},

		Route{
			"PostPost",
			strings.ToUpper("Post"),
			"/post",
			handler.PostPost,
		},

		Route{
			"PostUpdatePost",
			strings.ToUpper("Post"),
			"/post/update",
			handler.PostUpdatePost,
		},

		Route{
			"PostUploadContentPost",
			strings.ToUpper("Post"),
			"/post/upload/media/{postID}",
			handler.PostUploadMediaPost,
		},

		Route{
			"PostUploadContentPost",
			strings.ToUpper("Delete"),
			"/post/delete/media/{postID}",
			handler.PostDeleteMedia,
		},

		Route{
			"PostsPostIdDelete",
			strings.ToUpper("Delete"),
			"/delete/post/{postID}",
			handler.PostsPostIdDelete,
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

func handleOther(router *mux.Router, behavior interfaces.ContentBehavior) {
	op := "content.api.routers.NewRouterWithAuth"

	handler := api.New(behavior)

	var routes = Routes{
		Route{
			"FeedPopularGet",
			strings.ToUpper("Get"),
			"/feed/popular",
			handler.FeedPopularGet,
		},

		Route{
			"FeedSubscriptionsGet",
			strings.ToUpper("Get"),
			"/feed/subscriptions",
			handler.FeedSubscriptionsGet,
		},

		Route{
			"AuthorPostAuthorIdGet",
			strings.ToUpper("Get"),
			"/author/post/{authorID}",
			handler.AuthorPostAuthorIdGet,
		},

		Route{
			"PostMediaGet",
			strings.ToUpper("Get"),
			"/post/media/{postID}",
			handler.PostMediaGet,
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
}
