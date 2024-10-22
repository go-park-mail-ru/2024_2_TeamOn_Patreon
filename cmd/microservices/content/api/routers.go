package api

import (
	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/interfaces"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(behavior interfaces.ContentBehavior) *mux.Router {
	op := "content.api.routers"

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
			"/author/post/{authorId}",
			handler.AuthorPostAuthorIdGet,
		},

		Route{
			"PostLikePost",
			strings.ToUpper("Post"),
			"/post/like",
			handler.PostLikePost,
		},

		Route{
			"PostMediaGet",
			strings.ToUpper("Get"),
			"/post/media",
			handler.PostMediaGet,
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
			"/post/upload/content",
			handler.PostUploadContentPost,
		},

		Route{
			"PostsPostIdDelete",
			strings.ToUpper("Delete"),
			"/posts/{postId}",
			handler.PostsPostIdDelete,
		},
	}

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		logger.StandardInfoF(op, "Registered: %s %s", route.Method, route.Pattern)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
