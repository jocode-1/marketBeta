package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jocode-1/marketBeta/config"
	"github.com/jocode-1/marketBeta/internal/bootstrap"
	"github.com/jocode-1/marketBeta/internal/middleware"
)

// SetupRoutes initializes API routes
func SetupRoutes(router *gin.Engine, deps *bootstrap.AppDependencies) {
	router.Use(config.RequestMetricsMiddleware())
	api := router.Group("/api")
	{
		api.POST("/register", deps.UserHandler.Register)
		api.POST("/login", deps.UserHandler.Login)
	}

	protected := router.Group("/api/users").Use(middleware.AuthMiddleware())
	{
		protected.GET("/user/:user_id", deps.UserHandler.GetUserByID)

	}

	vendors := router.Group("/api/vendors").Use(middleware.AuthMiddleware())
	{
		vendors.POST("/createVendor", deps.VendorHandler.CreateVendor)

	}
}
