package helper

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ValidateBody[T any](r *http.Request) (T, error) {
	var data T

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, fmt.Errorf("invalid JSON: %w", err)
	}

	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return data, fmt.Errorf("validation failed: %w", err)
	}

	return data, nil
}
