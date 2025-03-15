package unit

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/jocode-1/marketBeta/internal/models"
	"github.com/jocode-1/marketBeta/internal/repositories"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
	"time"
)

func TestGetUserById(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	// Wrap sqlmock DB with sqlx
	sqlxDB := sqlx.NewDb(db, "postgres")

	// Create user repository instance
	userRepo := repositories.NewUserRepository(sqlxDB)

	// Sample UUID and expected user
	userID, err := uuid.Parse("550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		t.Fatalf("Failed to parse user UUID: %v", err)
	}
	expectedUser := &models.UserModel{
		UserID:          userID,
		UserName:        "John Doe",
		UserEmail:       "johndoe@example.com",
		HashedPassword:  "hashedpassword123",
		PhoneNumber:     "1234567890",
		UserAddress:     "123 Street, City",
		ProfilePhotoUrl: "Url",
		IpAddress:       "192.168.1.1",
		IsVerified:      true,
		IsAdmin:         false,
		IsVendor:        true,
		Role:            "vendor",
		Status:          true,
		UpdatedAt:       time.Now(),
		CreatedAt:       time.Now(),
	}

	// Mock database query response
	rows := sqlmock.NewRows([]string{
		"user_id", "username", "email", "hashed_password", "phone_number", "user_address", "profile_photo_url",
		"ip_address", "is_verified", "is_admin", "is_vendor",
		"role", "status", "updated_at", "created_at",
	}).AddRow(
		expectedUser.UserID, expectedUser.UserName, expectedUser.UserEmail, expectedUser.HashedPassword,
		expectedUser.PhoneNumber, expectedUser.UserAddress, expectedUser.ProfilePhotoUrl, expectedUser.IpAddress,
		expectedUser.IsVerified, expectedUser.IsAdmin, expectedUser.IsVendor, expectedUser.Role,
		expectedUser.Status, expectedUser.UpdatedAt, expectedUser.CreatedAt,
	)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT user_id, username, email, hashed_password, phone_number, user_address, profile_photo_url, ip_address, is_verified, is_admin, is_vendor, role, status, updated_at, created_at FROM users WHERE user_id = $1 LIMIT 1`)).
		WithArgs(userID).
		WillReturnRows(rows)

	// Call GetUserById
	result, err := userRepo.GetUserById(context.Background(), userID.String())

	// requireions
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, expectedUser.UserID, result.UserID)
	require.Equal(t, expectedUser.UserName, result.UserName)
	require.Equal(t, expectedUser.UserEmail, result.UserEmail)
	require.Equal(t, expectedUser.IsAdmin, result.IsAdmin)
	require.Equal(t, expectedUser.IsVendor, result.IsVendor)
	require.Equal(t, expectedUser.Role, result.Role)

	// Ensure all expectations were met
	require.NoError(t, mock.ExpectationsWereMet())
}
