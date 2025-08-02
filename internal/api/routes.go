package api

import (
	"fmt"
	"net/http"

	"github.com/parth469/go-basic-server/internal/api/middleware"
	v0 "github.com/parth469/go-basic-server/internal/api/v0"
)

type middlewareT func(http.Handler) http.Handler

func chainMiddleware(h http.Handler, m ...middlewareT) http.Handler {
	if len(m) == 0 {
		return h
	}

	wrapped := h
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}

	return wrapped
}

func APIRoutes() http.Handler {
	mux := http.NewServeMux()

	// Simple ping endpoint
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Pong")
	})

	commonMiddlewares := []middlewareT{
		middleware.CORSMiddleware,
		middleware.PanicRecovery,
		middleware.LoggingMiddleware,
		middleware.Auth,
	}

	mux.Handle("/", v0.Routes())

	wrappedHandler := chainMiddleware(mux, commonMiddlewares...)

	return http.StripPrefix("/api", wrappedHandler)
}
