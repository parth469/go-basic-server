package user

import (
	"fmt"
	"net/http"
	"time"

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

func Login(w http.ResponseWriter, r *http.Request) {
	user := UserRow{
		ID:        1,
		Name:      "John Doe",
		Email:     "john.doe@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  true,
	}

	token, err := helper.CreateToken(user)

	if err != nil {
		helper.ErrorWriter(w, r, 500, err)
		return
	}

	helper.ResponseWriter(w, r, map[string]interface{}{
		"token": token,
	})
	return

}
