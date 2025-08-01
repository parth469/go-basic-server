package database

import (
	"fmt"
	"time"

	"github.com/parth469/go-basic-server/util/config"
	"github.com/parth469/go-basic-server/util/logger"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const maxRetries = 3
const retryDelay = 2 * time.Second

var DB *gorm.DB

func Connect() error {
	dbURL := config.App.Database
	var err error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

		if err == nil {
			logger.Log.Info("Database connected successfully")
			break
		}

		logger.Log.Warn("Database connection attempt %d/%d failed: %v", attempt, maxRetries, err)
		if attempt < maxRetries {
			time.Sleep(retryDelay * time.Duration(attempt))
		}
	}

	if err != nil {
		return fmt.Errorf("failed to connect to database after %d retries: %w", maxRetries, err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get db instance: %w", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	sqlDB.SetConnMaxIdleTime(3 * time.Minute)

	return nil
}

func Close() error {
	sqlDB, err := DB.DB()
	if err != nil {
		logger.Log.Error("failed to retrieve database instance: %v", err)
		return err
	}

	if err := sqlDB.Close(); err != nil {
		logger.Log.Fatal("failed to close database connection: %v", err)
	} else {
		logger.Log.Info("Database connection closed successfully")
	}

	return nil
}

func Migrate() error {
	sqlDB, err := DB.DB()

	if err != nil {
		return fmt.Errorf("failed to retrieve database instance: %v", err)

	}

	migrationsDir := "internal/database/migrations"

	if err := goose.Up(sqlDB, migrationsDir); err != nil {
		logger.Log.Error("failed to apply migrations from '%s': %v", err)
		return err
	}

	logger.Log.Info("Database migrations applied successfully from '%s'", migrationsDir)
	return nil
}

func HealthCheck() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get db instance: %w", err)
	}

	return sqlDB.Ping()
}
