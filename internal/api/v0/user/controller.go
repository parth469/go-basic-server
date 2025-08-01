package user

import (
	"fmt"
	"net/http"

	"github.com/parth469/go-basic-server/util/helper"
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
	body, validateErr := helper.ValidateBody[CreateUser](r)

	if validateErr != nil {
		helper.ErrorWriter(w, r, 422, validateErr)
		return
	}

	creationErr := ProgressCreation(body)

	if creationErr != nil {
		helper.ErrorWriter(w, r, 500, creationErr)
		return

	}

	helper.ResponseWriter(w, r, body)

}
