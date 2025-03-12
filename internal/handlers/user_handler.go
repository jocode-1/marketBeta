package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jocode-1/marketBeta/config"
	"github.com/jocode-1/marketBeta/internal/models"
	"github.com/jocode-1/marketBeta/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

// Register a new user
func Register(c *gin.Context) {
	var user models.UserModel

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate fields
	if valid, msg := utils.ValidateUserInput(&user); !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.HashedPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.HashedPassword = string(hashedPassword)

	_, err = config.DB.NamedExec("INSERT INTO users (username, fullname, user_email, user_password, user_phone, user_address) VALUES (:username, :fullname, :user_email, :user_password, :user_phone, :user_address )", &user)

	if err != nil {
		log.Println("Database error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User already exists"})
		return
	}

	// Return response
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}
