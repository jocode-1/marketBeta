package bootstrap

import (
	"github.com/jmoiron/sqlx"
	"github.com/jocode-1/marketBeta/config"
	"github.com/jocode-1/marketBeta/internal/handlers"
	"github.com/jocode-1/marketBeta/internal/repositories"
	"github.com/sirupsen/logrus"
)

// AppDependencies holds all the dependencies for the application.
type AppDependencies struct {
	DB            *sqlx.DB
	Logger        *logrus.Logger
	UserHandler   *handlers.UserHandler
	VendorHandler *handlers.VendorHandler
}

// InitializeDependencies wires up all the dependencies.
func InitializeDependencies() *AppDependencies {
	// Connect to the database.
	db := config.ConnectDB() // returns *sqlx.DB

	// Get the centralized logger.
	logger := config.Logger

	// Initialize repository and handler using dependency injection.
	userRepo := repositories.NewUserRepository(db)
	vendorRepo := repositories.NewVendorRepository(db)
	userHandler := handlers.NewUserHandler(userRepo, logger)
	vendorHandler := handlers.NewVendorHandler(vendorRepo, logger)

	return &AppDependencies{
		DB:            db,
		Logger:        logger,
		UserHandler:   userHandler,
		VendorHandler: vendorHandler,
	}
}
