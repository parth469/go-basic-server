package api

import (
	"fmt"
	"net/http"

	v0 "github.com/parth469/go-basic-server/internal/api/v0"
)

func SetupAPIRoutes(route *http.ServeMux, prefix string) {

	route.HandleFunc(prefix+"/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Pong")
	})

	v0.V0Routes(route, prefix+"/v0")
}
