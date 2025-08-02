package middleware

import (
	"fmt"
	"net/http"

	"github.com/parth469/go-basic-server/util/helper"
)

var PublicRoutes = []string{
	"/v0/login",
	"/ping",
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		isPublic := func() bool {
			for _, publicRoute := range PublicRoutes {
				if path == publicRoute {
					return true
				}
			}
			return false
		}()

		if isPublic {
			next.ServeHTTP(w, r)
			return
		}

		helper.ErrorWriter(w, r, 401, fmt.Errorf("Please Provide Token"))
		return
	})
}
