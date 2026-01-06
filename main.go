package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

const PORT = "3030"

type apiConfig struct {
	fileserverHits atomic.Int32
}

type FileHandler struct{}

func (FileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fs := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
	fs.ServeHTTP(w, r)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now().Format("02/01/2006 15:04")
		fmt.Printf("[%s] %s %s\n", t, r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) metricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hits: %d", cfg.fileserverHits.Load())
}

func (cfg *apiConfig) resetHandler(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits.Store(0)
	w.WriteHeader(http.StatusOK)
}

func main() {
	mux := http.NewServeMux()
	apiCfg := apiConfig{}

	mux.HandleFunc("/healthz", HealthCheck)
	mux.HandleFunc("/metrics", apiCfg.metricsHandler)
	mux.HandleFunc("/reset", apiCfg.resetHandler)

	mux.Handle(
		"/app/",
		loggingMiddleware(
			apiCfg.middlewareMetricsInc(FileHandler{}),
		),
	)

	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
