package user

import (
	"net/http"
)

func UserRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /login", Login) // Create new user
	mux.HandleFunc("GET /users", Get)     // List all users
	mux.HandleFunc("POST /users", Create) // Create new user

	mux.HandleFunc("GET /users/{id}", GetById)   // Get specific user
	mux.HandleFunc("DELETE /users/{id}", Delete) // DELETE User
	mux.HandleFunc("PUT /users/{id}", Update)    // Update user

	return mux
}
