package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/parth469/go-basic-server/util/logger"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		recorder := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(recorder, r)

		duration := time.Since(start)

		status := recorder.status

		logMessage := fmt.Sprintf("Request | %s - %s | Status: %-3d | Duration: %d ms", r.Method, r.URL.Path, status, duration.Milliseconds())

		switch status {
		case http.StatusOK, http.StatusCreated:
			logger.Log.Info(logMessage)
		case http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusMethodNotAllowed, http.StatusConflict, http.StatusUnprocessableEntity, http.StatusTooManyRequests:
			logger.Log.Warn(logMessage)
		case http.StatusInternalServerError, http.StatusNotImplemented, http.StatusServiceUnavailable:
			logger.Log.Error(logMessage, nil)
		default:
			logger.Log.Info(logMessage)
		}
	})
}
