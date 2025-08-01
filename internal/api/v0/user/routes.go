package user

import (
	"net/http"
)

func UserRoutes(route *http.ServeMux, prefix string) {
	route.HandleFunc("GET "+prefix, Get)
}
