package repo

import (
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/utils/db"
	"log"
)

// SignInRepository implements SignInRepo interface for verifying user sign-in credentials.
type SignInRepository struct{}

// NewSignInRepository initializes and returns a new instance of SignInRepository.
// @Summary Initialize a new SignInRepository
// @Description Initializes and returns a new instance of SignInRepository.
// @return *SignInRepository
func NewSignInRepository() *SignInRepository {
	return &SignInRepository{}
}

// VerifySignInCredentials verifies the user's sign-in credentials against the database.
// @Summary Verify user sign-in credentials against the database
// @Description Verifies the user's sign-in credentials against the database.
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Param password path string true "Password"
// @Success 200 {string} string "true"
// @Failure 400 {string} string "false"
// @Router /verify [post]
func (r *SignInRepository) VerifySignInCredentials(username, password string) bool {
	var user models.SignInCredentials
	result := db.GormDb.Where("name = ?", username).First(&user)
	if result.Error != nil {
		log.Println("Error verifying user credentials:", result.Error)
		return false
	}

	// Compare hashed password with provided password
	if user.Password == password {
		return true // Passwords match (authentication successful)
	} else {
		return false // Passwords do not match (authentication failed)
	}
}
