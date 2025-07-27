package config

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Env  string `validate:"required,oneof=development production test"`
	Port string `json:"port" validate:"required"`
}

var App AppConfig

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Warning: .env file not found, using system environment variables")
	}

	App = AppConfig{
		Port: os.Getenv("PORT"),
		Env:  os.Getenv("ENVIRONMENT"),
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(App); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldErr := range validationErrors {
				log.Printf("Validation error: Field '%s' failed on the '%s' tag", fieldErr.Field(), fieldErr.Tag())
			}
			log.Fatal("Environment variable validation failed")
		} else {
			log.Fatalf("Unexpected validation error: %v", err)
		}
	}

	log.Println("Environment variable Load Successfully")
}
