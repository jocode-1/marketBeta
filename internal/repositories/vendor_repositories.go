package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/jocode-1/marketBeta/internal/models"
)

type VendorRepository interface {
	CreateVendor(ctx context.Context, vendor *models.VendorProfile) error
}

type vendorRepository struct {
	db *sqlx.DB
}

func NewVendorRepository(db *sqlx.DB) VendorRepository {
	return &vendorRepository{db: db}
}

func (r *vendorRepository) CreateVendor(ctx context.Context, vendor *models.VendorProfile) error {
	query := `INSERT INTO vendor_profiles (user_id, business_id, business_name, category, address, website, cac_number, payment_made, is_verified, is_business_verified, 
                    verification_status, rejected_reason, total_sales, rating, review_count, status, created_at, updated_at) VALUES ($1, $2, $3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18) RETURNING 
                    user_id, business_id, business_name, category, address, website, cac_number, payment_made, is_verified, is_business_verified, verification_status, rejected_reasons, total_sales, rating, review_count, status, created_at, updated_at`

	return r.db.QueryRowContext(ctx, query, vendor.UserID, vendor.BusinessID, vendor.BusinessName, vendor.Category, vendor.Address, vendor.Website, vendor.CacNumber, vendor.PaymentMade, vendor.Is_verified,
		vendor.Is_business_verified, vendor.VerificationStatus, vendor.RejectedReason, vendor.TotalSales, vendor.Rating, vendor.ReviewsCount, vendor.Status, vendor.CreatedAt, vendor.UpdatedAt).Scan(&vendor.UserID, &vendor.BusinessID,
		&vendor.BusinessName, &vendor.Category, &vendor.Address, &vendor.Website, &vendor.CacNumber, &vendor.PaymentMade, &vendor.Is_verified,
		&vendor.Is_business_verified, &vendor.VerificationStatus, &vendor.RejectedReason, &vendor.TotalSales, &vendor.Rating, &vendor.ReviewsCount, &vendor.Status, &vendor.CreatedAt, &vendor.UpdatedAt)
}
