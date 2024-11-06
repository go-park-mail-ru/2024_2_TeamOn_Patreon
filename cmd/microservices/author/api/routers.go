package api

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares"

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

func NewRouter(service interfaces.AuthorService) *mux.Router {
	op := "author.routers.NewRouter"

	handler := api.New(service)

	var routes = Routes{
		Route{
			"GetAuthorPayments",
			"GET",
			"/author/payments",
			handler.GetAuthorPayments,
		},
		Route{
			"GetAuthorBackground",
			"GET",
			"/author/{authorID}/background",
			handler.GetAuthorBackground,
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
			"GetCSRFToken",
			strings.ToUpper("Post"),
			"/author/{authorId}/following",
			handler.PostFollowing,
		},
	}
	// Declare a new router
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		logger.StandardInfo(
			context.Background(),
			fmt.Sprintf("Registered: %s %s", route.Method, route.Pattern),
			op,
		)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
