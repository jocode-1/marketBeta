package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jocode-1/marketBeta/config"
	"github.com/jocode-1/marketBeta/internal/bootstrap"
	"github.com/jocode-1/marketBeta/internal/middleware"
	"github.com/jocode-1/marketBeta/routes"
)

// SetupRouter initializes the Gin router without starting the server
func SetupRouter() *gin.Engine {
	config.InitLogger()                                   // ✅ Initialize logging first
	config.Logger.Info("Logger initialized successfully") // ✅ Confirm it works

	config.ConnectDB()      // ✅ Connect to DB
	config.InitMonitoring() // ✅ Setup monitoring

	deps := bootstrap.InitializeDependencies() // ✅ Inject dependencies

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.RateLimitMiddleware())

	routes.SetupRoutes(router, deps) // ✅ Pass dependencies to routes
	return router
}

// StartServer starts the HTTP server
func StartServer() {
	router := SetupRouter()
	config.Logger.Info("Starting eCommerce API on port 8080...") // ✅ Log before starting

	if err := router.Run(":8080"); err != nil {
		config.Logger.Fatal("Failed to start server:", err) // ✅ Log fatal error
	}
}

func main() {
	StartServer() // ✅ Start the server properly
}
