package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/parth469/go-basic-server/internal/api/v0/user"
	"github.com/parth469/go-basic-server/util/helper"
)

const (
	AuthorizationHeader = "Authorization"
	BearerPrefix        = "Bearer "
)

var PublicRoutes = map[string]bool{
	"/v0/login": true,
	"/ping":     true,
}

type contextKey string

const (
	UserID contextKey = "userId"
)

func isPublicRoute(path string) bool {
	return PublicRoutes[path]
}

func extractToken(header string) string {
	if strings.HasPrefix(header, BearerPrefix) {
		return strings.TrimSpace(header[len(BearerPrefix):])
	}
	return strings.TrimSpace(header)
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isPublicRoute(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get(AuthorizationHeader)
		if authHeader == "" {
			w.Header().Set("WWW-Authenticate", "Bearer")
			helper.ErrorWriter(w, r, http.StatusUnauthorized, fmt.Errorf("authorization token is missing"))
			return
		}

		payload, err := helper.VerifyToken(extractToken(authHeader))
		if err != nil {
			helper.ErrorWriter(w, r, http.StatusForbidden, fmt.Errorf("authorization token verification failed: %v", err))
			return
		}

		userID, ok := payload["id"]
		if !ok {
			helper.ErrorWriter(w, r, http.StatusForbidden, fmt.Errorf("userId not found in token payload"))
			return
		}
		userIDStr := fmt.Sprintf("%v", userID)
		userRecord, err := user.FetchUserByID(userIDStr)

		ctx := context.WithValue(r.Context(), UserID, userRecord)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
