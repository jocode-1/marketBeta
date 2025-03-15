package models

import (
	"github.com/google/uuid"
	"time"
)

type VendorProfile struct {
	ID           uuid.UUID `db:"id" json:"id"`
	UserID       uuid.UUID `db:"user_id" json:"user_id"`
	BusinessID   uuid.UUID `db:"business_id" json:"business_id"`
	BusinessName string    `db:"business_name" json:"business_name"`
	Category     string    `db:"category" json:"category"`
	Address      string    `db:"address" json:"address"`
	Website      string    `db:"website" json:"website,omitempty"`
	TaxID        string    `db:"tax_id" json:"tax_id,omitempty"`
	PaymentMade  bool      `db:"payment_made" json:"payment_made"`
	IsVerified   bool      `db:"verified" json:"verified"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}
