package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jocode-1/marketBeta/internal/handlers"
	"github.com/jocode-1/marketBeta/internal/middleware"
)

// SetupRoutes initializes API routes
func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)
		//api.GET("/refresh-token", controllers.RefreshToken)
	}

	protected := router.Group("/api/protected").Use(middleware.AuthMiddleware())
	{
		protected.GET("/dashboard", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome to your dashboard"})
		})
	}
}
