package models

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

// VendorProfile struct represents the vendor's business profile.
type VendorProfile struct {
	ID                   uuid.UUID      `db:"id" json:"id"`
	UserID               uuid.UUID      `db:"user_id" json:"user_id"`
	BusinessID           uuid.UUID      `db:"business_id" json:"business_id"`
	BusinessName         string         `db:"business_name" json:"business_name"`
	Category             string         `db:"category" json:"category"`
	Address              string         `db:"address" json:"address"`
	Website              string         `db:"website" json:"website,omitempty"`
	CacNumber            string         `db:"cac_number" json:"cac_number,omitempty"`
	PaymentMade          bool           `db:"payment_made" json:"payment_made"`
	Is_verified          bool           `db:"is_verified" json:"is_verified"`
	Is_business_verified bool           `db:"is_business_verified" json:"is_business_verified"`
	VerificationStatus   string         `db:"verification_status" json:"verification_status"`
	RejectedReason       sql.NullString `db:"rejected_reason" json:"rejected_reason,omitempty"`

	TotalSales   float64 `db:"total_sales" json:"total_sales,omitempty"`
	Rating       float64 `db:"rating" json:"rating,omitempty"`
	ReviewsCount int     `db:"reviews_count" json:"reviews_count,omitempty"`

	Status    string    `db:"status" json:"status"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
