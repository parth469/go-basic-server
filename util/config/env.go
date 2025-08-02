package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/parth469/go-basic-server/util/logger"
	"os"
)

type AppConfig struct {
	Env       string `validate:"required,oneof=development production test"`
	Port      string `json:"port" validate:"required"`
	Database  string `json:"database" validate:"required"`
	SecretKey string `json:"secret_Key" validate:"required"`
}

var App AppConfig

func Load() {
	if err := godotenv.Load(); err != nil {
		logger.Log.Fatal("Warning: .env file not found, using system environment variables", err)
	}

	App = AppConfig{
		Port:      os.Getenv("PORT"),
		Env:       os.Getenv("ENVIRONMENT"),
		Database:  os.Getenv("DATABASE"),
		SecretKey: os.Getenv("SECRETKEY"),
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(App); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldErr := range validationErrors {
				logger.Log.Warn("Validation error: Field '%s' failed on the '%s' tag", fieldErr.Field(), fieldErr.Tag())
			}
			logger.Log.Fatal("Environment variable validation failed", nil)
		} else {
			logger.Log.Fatal("Unexpected validation error: %v", err)
		}
	}

	logger.Log.Info("Environment variable Load Successfully")
}
