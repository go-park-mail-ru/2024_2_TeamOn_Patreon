package api

import (
	"encoding/json"
	"fmt"
	models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/polina-auth/internal/auth/api/models"
	"log/slog"
	"net/http"
)

// AuthLoginPost - ручка аутентификации
func AuthLoginPost(w http.ResponseWriter, r *http.Request) {
	op := "auth.api.api_auth.AuthLoginPost"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// parse models
	var l models.Login
	if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
		slog.Info(fmt.Sprintf("error {%v} | in %v", err, op))
		// TODO: Дописать отправку модели ошибки с err.msg
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// validate fields
	if _, errV := l.Validate(); errV != nil {
		slog.Info(fmt.Sprintf("error {%v} | in %v %v", errV, op, errV.GetMsg()))

		// TODO: Дописать отправку модели ошибки с err.msg
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	slog.Info(fmt.Sprintf("valid request | in %v", op))
	w.WriteHeader(http.StatusOK)
}

// AuthRegisterPost - ручка регистрации
func AuthRegisterPost(w http.ResponseWriter, r *http.Request) {
	op := "auth.api.api_auth.AuthRegisterPost"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// parse models
	var p models.Reg
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		slog.Info(fmt.Sprintf("error {%v} | in %v", err, op))
		// TODO: Дописать отправку модели ошибки с err.msg
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// validate fields
	if _, errV := p.Validate(); errV != nil {
		slog.Info(fmt.Sprintf("error {%v} | in %v %v", errV, op, errV.GetMsg()))

		// TODO: Дописать отправку модели ошибки с err.msg
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	slog.Info(fmt.Sprintf("valid request | in %v", op))
	w.WriteHeader(http.StatusOK)
}
