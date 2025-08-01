package main

import (
	"net/http"

	"github.com/parth469/go-basic-server/internal/api"
	"github.com/parth469/go-basic-server/internal/database"
	"github.com/parth469/go-basic-server/util/config"
	"github.com/parth469/go-basic-server/util/logger"
)

func main() {
	config.Load()

	database.Connect()
	defer database.Close()

	database.Migrate()

	route := http.NewServeMux()

	api.SetupAPIRoutes(route, "/api")

	server := http.Server{
		Addr:    config.App.Port,
		Handler: route,
	}

	logger.Log.Info("Starting Server...")

	if err := server.ListenAndServe(); err != nil {
		logger.Log.Fatal("Server Fail to Start", err)
	}
}
