package v0

import (
	"net/http"

	"github.com/parth469/go-basic-server/internal/api/v0/user"
)

func Routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", user.UserRoutes())

	return http.StripPrefix("/v0", mux)
}
