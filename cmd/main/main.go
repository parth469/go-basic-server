package main

import (
	"fmt"
	"net/http"
)

func main() {
	route := http.NewServeMux()

	route.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to basic golang server")
	})

	server := http.Server{
		Addr:    ":3030",
		Handler: route,
	}

	server.ListenAndServe()
}
