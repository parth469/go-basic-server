package main

import (
	"fmt"
	"log"
	"net/http"

	config "github.com/parth469/go-basic-server/util/config"
)

func main() {
	config.Load()
	route := http.NewServeMux()

	route.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to basic golang server")
	})

	server := http.Server{
		Addr:    config.App.Port,
		Handler: route,
	}

	fmt.Println("Server Start on", server.Addr)

	log.Fatal(server.ListenAndServe())
}
