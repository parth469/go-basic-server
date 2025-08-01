package user

import (
	"fmt"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GET USER")
}

func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "UPDATE USER")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "DELETE USER")
}

func GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GET USER BY ID")
}

func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "CREATE USER")
}
