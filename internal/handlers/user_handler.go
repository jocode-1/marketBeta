package handlers

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	v "github.com/go-ozzo/ozzo-validation/v4"
	is "github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/jocode-1/marketBeta/config"
	"github.com/jocode-1/marketBeta/internal/models"
	"github.com/jocode-1/marketBeta/internal/repositories"
	"github.com/jocode-1/marketBeta/internal/utils"
	"github.com/jocode-1/marketBeta/queries"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

type UserHandler struct {
	repo   repositories.UserRepository
	logger *logrus.Logger
}

// NewUserHandler creates a new UserHandler with its dependencies injected.
func NewUserHandler(repo repositories.UserRepository, logger *logrus.Logger) *UserHandler {
	return &UserHandler{
		repo:   repo,
		logger: logger,
	}
}

// SignupUserParam defines the expected input for user registration
type SignupUserParam struct {
	Username        string `json:"username"`
	Email           string `json:"user_email"`
	Password        string `json:"password"`
	PhoneNumber     string `json:"phone_number"`
	UserAddress     string `json:"address,omitempty"`
	ProfilePhotoUrl string `json:"profile_photo_url,omitempty"`
}

// validate validates the SignupUserParam fields
func (s SignupUserParam) validate() error {
	return v.ValidateStruct(&s,
		v.Field(&s.Username, v.Skip),
		v.Field(&s.Email, v.Required, is.Email),
		v.Field(&s.Password, v.Required, v.Length(6, 100)),
		v.Field(&s.PhoneNumber, v.Required, v.Length(12, 100)),
		v.Field(&s.UserAddress, v.Required, v.Length(6, 100)),
		v.Field(&s.ProfilePhotoUrl, v.Skip),
	)
}

// Register a new user
func (h *UserHandler) Register(c *gin.Context) {
	var signupParam SignupUserParam

	if err := c.ShouldBindJSON(&signupParam); err != nil {
		h.logger.Error("Invalid request payload")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate fields
	if err := signupParam.validate(); err != nil {
		h.logger.Error("User Validation failed: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupParam.Password), bcrypt.DefaultCost)
	if err != nil {
		h.logger.Error("Hashing password failed: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Get User's IP Address
	userIP := utils.GetIPAddress(c.Request)

	user := models.UserModel{
		UserID:         uuid.New(),
		UserName:       signupParam.Username,
		UserEmail:      signupParam.Email,
		HashedPassword: string(hashedPassword),
		UserAddress:    signupParam.UserAddress,
		PhoneNumber:    signupParam.PhoneNumber,
		IpAddress:      sql.NullString{String: userIP, Valid: userIP != ""},
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err = h.repo.CreateUser(context.Background(), &user); err != nil {
		h.logger.Error("Failed to create user: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User already exists or creation failed"})
		return
	}

	// Return response
	h.logger.Info("User registered successfully: ", user.UserEmail)
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}

// Login authenticates user

type LoginUserParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u LoginUserParam) validate() error {
	return v.ValidateStruct(&u,
		v.Field(&u.Email, v.Required, is.Email),
		v.Field(&u.Password, v.Required),
	)
}

func Login(c *gin.Context) {
	var input LoginUserParam

	// Bind and validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := input.validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Define user model to hold fetched data
	var user models.UserModel

	// Fetch user from database using sqlx
	err := config.DB.Get(&user, queries.GetUserByEmail, input.Email)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	} else if err != nil {
		log.Println("Database error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Generate JWT tokens
	accessToken, refreshToken, err := utils.GenerateTokens(user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"user_id":   user.UserID,
			"username":  user.UserName,
			"email":     user.UserEmail,
			"is_admin":  user.IsAdmin,
			"is_vendor": user.IsVendor,
			"role":      user.Role,
		},
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// RefreshToken handles token renewal using the refresh token
func RefreshToken(c *gin.Context) {
	var request struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	claims, err := utils.ValidateRefreshToken(request.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	userIDStr, ok := claims["user_id"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID format"})
		return
	}

	userUUID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	newAccessToken, newRefreshToken, err := utils.GenerateTokens(userUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": newAccessToken, "refresh_token": newRefreshToken})
}
