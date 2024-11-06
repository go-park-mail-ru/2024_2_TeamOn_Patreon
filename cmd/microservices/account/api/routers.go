package api

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares"

	"net/http"

	api "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/controller"
	interfaces "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/controller/interfaces"
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

func NewRouter(service interfaces.AccountService) *mux.Router {
	op := "account.routers.NewRouter"

	handler := api.New(service)

	var routes = Routes{
		Route{
			"GetAccount",
			"GET",
			"/account",
			handler.GetAccount,
		},
		Route{
			"GetAccountAvatar",
			"GET",
			"/account/{userID}/avatar",
			handler.GetAccountAvatar,
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
			"GetCSRFToken",
			strings.ToUpper("Get"),
			"/token-endpoint",
			middlewares.GetCSRFTokenHandler,
		},
	}
	// Declare a new router
	router := mux.NewRouter().StrictSlash(true)

	ctx := context.Background()

	for _, route := range routes {
		logger.StandardInfo(
			ctx,
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
