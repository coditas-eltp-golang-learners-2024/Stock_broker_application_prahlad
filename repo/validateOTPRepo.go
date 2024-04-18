package repo

import (
	"Stock_broker_application/src/app/authentication/models"
	"log"
	"time"

	"gorm.io/gorm"
)

// OTPValidationRepo defines the interface for OTP-related database operations.
type OTPValidationRepo interface {
	InsertOTP(otp int, createdAt int64, email string) error
	GetOTP(email string) (models.OTP, error)
}

// DefaultOTPRepo implements OTPValidationRepo using GORM and a database connection.
type DefaultOTPValidationRepo struct {
	DB *gorm.DB
}

// GetOTP retrieves an OTP record from the database based on the provided email.
func (repo *DefaultOTPValidationRepo) GetOTP(email string) (models.OTP, error) {
	var OTP models.OTP

	if err := repo.DB.Table("users").Where("email = ?", email).First(&OTP).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.OTP{}, err
		}
		return models.OTP{}, err
	}
	return OTP, nil
}

// InsertOTP inserts or updates an OTP record in the database for the given email.
func (repo *DefaultOTPValidationRepo) InsertOTP(otp int, createdAt int64, Email string) error {
	if repo.DB == nil {
		log.Fatal("Database connection is not initialized")
	}

	// Update the OTP column for the given email
	if err := repo.DB.Model(&models.OTP{}).Where("email = ?", Email).Update("otp", otp).Error; err != nil {
		return err
	}

	// Update the createdAt column for the given email
	currentTime := time.Now()
	if err := repo.DB.Model(&models.OTP{}).Where("email = ?", Email).Update("createdAt", currentTime).Error; err != nil {
		return err
	}

	// Update the epochTimestamp column for the given email
	if err := repo.DB.Model(&models.OTP{}).Where("email = ?", Email).Update("epochTimestamp", createdAt).Error; err != nil {
		return err
	}

	return nil
}
