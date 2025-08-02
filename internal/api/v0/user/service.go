package user

import "time"

func ProgressCreation(body CreateUser) error {
	return nil
}

func FetchUserByID(id string) (UserRow, error) {
	user := UserRow{
		ID:        1,
		Name:      "John Doe",
		Email:     "john.doe@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  true,
	}

	return user, nil

}
