package middleware

import (
	"fmt"
	"net/http"

	"github.com/parth469/go-basic-server/util/logger"
)

func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				logger.Log.Error(fmt.Sprintf("panic recovered: %v", err), nil)
			}
		}()
		next.ServeHTTP(w, req)
	})
}
