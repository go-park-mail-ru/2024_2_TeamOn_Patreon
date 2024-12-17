package api

import (
	"context"
	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/interfaces"
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

func NewRouter(behavior interfaces.ContentBehavior, monster *middlewares.Monster) *mux.Router {
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

func handleAuth(router *mux.Router, behavior interfaces.ContentBehavior) *mux.Router {
	op := "content.api.routers.NewRouterWithAuth"

	handler := api.New(behavior)

	var routes = Routes{
		Route{
			"PostLikePost",
			http.MethodPost,
			"/post/like",
			handler.PostLikePost,
		},

		Route{
			"PostPost",
			http.MethodPost,
			"/post",
			handler.PostPost,
		},

		Route{
			"PostUpdatePost",
			http.MethodPost,
			"/post/update",
			handler.PostUpdatePost,
		},

		Route{
			"PostUploadContentPost",
			http.MethodPost,
			"/post/upload/media/{postID}",
			handler.PostUploadMediaPost,
		},

		Route{
			"PostUploadContentPost",
			http.MethodDelete,
			"/post/delete/media/{postID}",
			handler.PostDeleteMedia,
		},

		Route{
			"PostsPostIdDelete",
			http.MethodDelete,
			"/delete/post/{postID}",
			handler.PostsPostIdDelete,
		},

		Route{
			"PostsCommentsCommentIDDeleteDelete",
			http.MethodDelete,
			"/posts/comments/{commentID}/delete",
			handler.PostsCommentsCommentIDDeleteDelete,
		},

		Route{
			"PostsCommentsCommentIDUpdatePost",
			http.MethodPost,
			"/posts/comments/{commentID}/update",
			handler.PostsCommentsCommentIDUpdatePost,
		},

		Route{
			"PostsPostIDCommentsCreatePost",
			http.MethodPost,
			"/posts/{postID}/comments/create",
			handler.PostsPostIDCommentsCreatePost,
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
			http.MethodGet,
			"/feed/popular",
			handler.FeedPopularGet,
		},

		Route{
			"FeedSubscriptionsGet",
			http.MethodGet,
			"/feed/subscriptions",
			handler.FeedSubscriptionsGet,
		},

		Route{
			"AuthorPostAuthorIdGet",
			http.MethodGet,
			"/author/post/{authorID}",
			handler.AuthorPostAuthorIdGet,
		},

		Route{
			"PostMediaGet",
			http.MethodGet,
			"/post/media/{postID}",
			handler.PostMediaGet,
		},
		Route{
			"PostsPostIDCommentsGet",
			http.MethodGet,
			"/posts/{postID}/comments",
			handler.PostsPostIDCommentsGet,
		},
		Route{
			"GetCSRFToken",
			http.MethodGet,
			"/token-endpoint",
			middlewares.GetCSRFTokenHandler,
		},
		Route{
			"Metrics",
			http.MethodGet,
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
