package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// RoleType defines the type of role (e.g., Admin, Vendor, User)
type RoleType string

// Role struct
type Role struct {
	ID        uuid.UUID `db:"id"`
	Name      RoleType  `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// RoleUser struct (Join table for many-to-many relationship)
type RoleUser struct {
	ID        uuid.UUID `db:"id"`
	UserID    uuid.UUID `db:"user_id"`
	RoleID    uuid.UUID `db:"role_id"`
	CreatedAt time.Time `db:"created_at"`
}

// UserModel struct
type UserModel struct {
	UserID          uuid.UUID      `db:"user_id" json:"user_id"`
	UserName        string         `db:"username" json:"username"`
	UserEmail       string         `db:"email" json:"email"`
	HashedPassword  string         `db:"hashed_password" json:"-"`
	PhoneNumber     string         `db:"phone_number" json:"phone_number"`
	UserAddress     string         `db:"user_address" json:"address"`
	ProfilePhotoUrl sql.NullString `db:"profile_photo_url" json:"profile_photo_url"`
	IpAddress       string         `db:"ip_address" json:"ip_address"`
	IsVerified      bool           `db:"is_verified" json:"is_verified"`
	IsAdmin         bool           `db:"is_admin" json:"is_admin"`
	IsVendor        bool           `db:"is_vendor" json:"is_vendor"`
	Role            string         `db:"role" json:"role"`
	Status          bool           `db:"status" json:"status"`
	UpdatedAt       time.Time      `db:"updated_at" json:"updated_at"`
	CreatedAt       time.Time      `db:"created_at" json:"created_at"`
}
