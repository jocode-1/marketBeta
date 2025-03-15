package repositories

import (
	"context"
	"github.com/jocode-1/marketBeta/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.UserModel) error
	GetUserByEmail(ctx context.Context, email string) (*models.UserModel, error)
	GetUserById(ctx context.Context, userID string) (*models.UserModel, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.UserModel) error {
	query := `
INSERT INTO users (
    username, email, hashed_password, phone_number, user_address, ip_address, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING user_id, username, email, hashed_password, phone_number, user_address, profile_photo_url,
          ip_address, is_verified, is_admin, is_vendor, role, status, updated_at, created_at
`

	return r.db.QueryRowContext(ctx, query,
		user.UserName,
		user.UserEmail,
		user.HashedPassword,
		user.PhoneNumber,
		user.UserAddress,
		user.IpAddress,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(
		&user.UserID,
		&user.UserName,
		&user.UserEmail,
		&user.HashedPassword,
		&user.PhoneNumber,
		&user.UserAddress,
		&user.ProfilePhotoUrl,
		&user.IpAddress,
		&user.IsVerified,
		&user.IsAdmin,
		&user.IsVendor,
		&user.Role,
		&user.Status,
		&user.UpdatedAt,
		&user.CreatedAt,
	)
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.UserModel, error) {
	var user models.UserModel
	query := `SELECT user_id, username, email, hashed_password, phone_number, user_address, profile_photo_url, 
       ip_address, is_verified, is_admin, is_vendor, role, status, updated_at, created_at FROM users WHERE email = $1 LIMIT 1`
	err := r.db.GetContext(ctx, &user, query, email)
	return &user, err
}

func (r *userRepository) GetUserById(ctx context.Context, userId string) (*models.UserModel, error) {
	var user models.UserModel
	query := `SELECT user_id, username, email, hashed_password, phone_number, user_address, profile_photo_url, 
       ip_address, is_verified, is_admin, is_vendor, role, status, updated_at, created_at FROM users WHERE user_id = $1 LIMIT 1`
	err := r.db.GetContext(ctx, &user, query, userId)
	return &user, err
}
