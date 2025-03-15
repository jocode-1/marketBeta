package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jocode-1/marketBeta/internal/models"
	"github.com/jocode-1/marketBeta/internal/utils"
	"net/http"
	"strings"
)

type contextKey string

const authUserCtxKey contextKey = "authUser"

// AuthMiddleware verifies JWT token in headers
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		token := tokenParts[1]
		claims, err := utils.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}

func GetAuthUserFromContext(ctx *gin.Context) (*models.UserModel, error) {
	// Retrieve the user payload from the request context
	payload, exists := ctx.Get(string(authUserCtxKey))
	if !exists {
		return nil, errors.New("unauthorized: no user found in context")
	}
	user, ok := payload.(*models.UserModel)
	if !ok || user == nil {
		return nil, errors.New("invalid user data in context")
	}

	return user, nil
}
