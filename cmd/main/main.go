package main

import (
	"fmt"
	"net/http"

	config "github.com/parth469/go-basic-server/util/config"
	"github.com/parth469/go-basic-server/util/logger"
)

func main() {
	config.Load()
	route := http.NewServeMux()

	route.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		logger.Log.Info("Hello")
		fmt.Fprintln(w, "test")
	})

	server := http.Server{
		Addr:    config.App.Port,
		Handler: route,
	}

	logger.Log.Info("Starting Server...")

	if err := server.ListenAndServe(); err != nil {
		logger.Log.Fatal("Server Fail to Start", err)
	}
}
