package controller

import "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/controller/interfaces"

type Handler struct {
	serv interfaces.AuthorService
}

func New(service interfaces.AuthorService) *Handler {
	return &Handler{serv: service}
}

const (
	// Path param
	authorIDParam string = "authorID"
	anon          string = "anonim"
)
