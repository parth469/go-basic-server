package v0

import (
	"net/http"

	"github.com/parth469/go-basic-server/internal/api/v0/user"
)

func V0Routes(route *http.ServeMux, prefix string) {
	user.UserRoutes(route, prefix+"/user")
}
