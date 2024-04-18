package service

import (
	"log"
	"math/rand"
	"time"

	"Stock_broker_application/src/app/authentication/repo"
)

// OTPValidationService defines the interface for OTP-related operations.
type OTPValidationService interface {
	GenerateOTP() (int, int64)
	ValidateOTP(otp int, email string, repo repo.OTPValidationRepo) (bool, error)
}

// DefaultOTPValidationService implements OTPValidationService for OTP generation and validation.
type DefaultOTPValidationService struct{}

// GenerateOTP generates a random 4-digit OTP and returns it along with the current epoch timestamp.
func (s *DefaultOTPValidationService) GenerateOTP() (int, int64) {
	source := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(source)

	// Generate a random 4-digit number between 1000 and 9999
	randomNumber := randomGenerator.Intn(9000) + 1000

	// Calculate the current epoch timestamp
	epochTimeStamp := time.Now().Unix()

	return randomNumber, epochTimeStamp
}

// ValidateOTP validates the provided OTP against the stored OTP for the given email.
func (serviceStruct *DefaultOTPValidationService) ValidateOTP(otp int, email string, repo repo.OTPValidationRepo) (bool, error) {
	// Fetch stored OTP details from the repository
	storedOTP, err := repo.GetOTP(email)
	if err != nil {
		return false, err
	}
	providedOTP := otp

	// Check if the provided OTP matches the stored OTP value
	if storedOTP.Value == providedOTP {
		// Calculate the current time
		currentTime := time.Now()

		// Calculate the time difference between OTP insertion time and current time
		otpInsertedTime := time.Unix(storedOTP.EpochTimeStamp, 0)
		timeDiff := currentTime.Sub(otpInsertedTime)

		// Check if the OTP is within the validity period (e.g., 5 minutes)
		if timeDiff <= 2*time.Minute {
			log.Println("Valid OTP")
			return true, nil
		} else {
			log.Println("OTP expired")
			return false, nil
		}
	}

	log.Println("Invalid OTP")
	return false, nil
}
