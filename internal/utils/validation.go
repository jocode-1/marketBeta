package utils

import (
	"github.com/jocode-1/marketBeta/internal/models"
	"regexp"
)

// Validate email format
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// Validate user input fields

func ValidateUserInput(user *models.UserModel) (bool, string) {
	if user.UserName == "" || len(user.UserName) < 3 {
		return false, "Username must be at least 3 characters"
	}
	if user.UserEmail == "" || !isValidEmail(user.UserEmail) {
		return false, "Invalid email format"
	}
	if user.HashedPassword == "" || len(user.HashedPassword) < 6 {
		return false, "Password must be at least 6 characters"
	}
	return true, ""
}
