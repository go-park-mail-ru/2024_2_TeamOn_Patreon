package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Main page")
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Авторизация:", r.URL.String())
}

func main() {
	http.HandleFunc("/auth/", authHandler)

	http.HandleFunc("/", handler)

	fmt.Println("starting server at: 8080")
	http.ListenAndServe(":8080", nil)
}
