package api

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/interfaces"
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

func NewRouter(behavior interfaces.ContentBehavior) *mux.Router {
	mainRouter := mux.NewRouter().StrictSlash(true)
	return mainRouter
}
