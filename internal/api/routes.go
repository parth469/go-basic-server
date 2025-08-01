package api

import (
	"fmt"
	"net/http"

	v0 "github.com/parth469/go-basic-server/internal/api/v0"
)


func APIRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Pong")
	})

	mux.Handle("/", v0.Routes())

	return http.StripPrefix("/api", mux)
}