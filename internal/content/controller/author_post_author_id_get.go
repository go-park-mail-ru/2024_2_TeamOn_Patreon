package controller

import "net/http"

func (h Handler) AuthorPostAuthorIdGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	// сделано на пк
}
