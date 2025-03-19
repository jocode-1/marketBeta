package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/jocode-1/marketBeta/internal/models"
	"github.com/jocode-1/marketBeta/internal/repositories"
	"github.com/sirupsen/logrus"
	"net/http"
)

type VendorHandler struct {
	repo   repositories.VendorRepository
	logger *logrus.Logger
}

func NewVendorHandler(repo repositories.VendorRepository, logger *logrus.Logger) *VendorHandler {
	return &VendorHandler{
		repo:   repo,
		logger: logger,
	}
}

type VendorProfileParam struct {
	BusinessName string `json:"business_name"`
	Category     string `json:"category"`
	Address      string `json:"address"`
	Website      string `json:"website,omitempty"`
	CacNumber    string `json:"cac_number,omitempty"`
}

func (vd VendorProfileParam) validate() error {
	return v.ValidateStruct(&vd,
		v.Field(&vd.BusinessName, v.Required),
		v.Field(&vd.Category, v.Required),
		v.Field(&vd.Address, v.Required),
		v.Field(&vd.Website, v.Skip),
		v.Field(&vd.CacNumber, v.Required),
	)
}

func (ve *VendorHandler) CreateVendor(c *gin.Context) {

	var vendorparam VendorProfileParam

	if err := c.ShouldBindJSON(&vendorparam); err != nil {
		ve.logger.Error("Invalid request payload")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := vendorparam.validate(); err != nil {
		ve.logger.Error("User Validation failed: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vendor := models.VendorProfile{
		UserID:       uuid.New(),
		BusinessID:   uuid.New(),
		BusinessName: vendorparam.BusinessName,
		Category:     vendorparam.Category,
		Address:      vendorparam.Address,
		Website:      vendorparam.Website,
		CacNumber:    vendorparam.CacNumber,
	}

	if err := ve.repo.CreateVendor(context.Background(), &vendor); err != nil {
		ve.logger.Error("Failed to create Vendor: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Vendor already exists or creation failed"})
		return
	}

	// Return response
	ve.logger.Info("User registered successfully: ", vendor)
	c.JSON(http.StatusCreated, gin.H{
		"message":       "Vendor registered successfully",
		"vendorProfile": vendor,
	})

}
