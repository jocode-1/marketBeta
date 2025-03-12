package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jocode-1/marketBeta/config"
	"github.com/jocode-1/marketBeta/internal/middleware"
	"github.com/jocode-1/marketBeta/routes"
)

// SetupRouter initializes the Gin router without starting the server
func SetupRouter() *gin.Engine {
	config.ConnectDB()

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.RateLimitMiddleware())

	routes.SetupRoutes(router)

	return router
}

// StartServer starts the HTTP server
func StartServer() {
	router := SetupRouter()
	router.Run(":8080") // Start server on port 8080
}

func main() {
	StartServer() // Manually start the server
}
