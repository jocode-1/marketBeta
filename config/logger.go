package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Global logger instance
var Logger *logrus.Logger

// InitLogger initializes the Logrus logger
func InitLogger() {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.JSONFormatter{}) // JSON formatting for structured logs
	Logger.SetOutput(os.Stdout)                  // Print logs to console
	Logger.SetLevel(logrus.InfoLevel)            // Set log level to INFO

	Logger.Info("Logger initialized successfully") // Ensure initialization
}
