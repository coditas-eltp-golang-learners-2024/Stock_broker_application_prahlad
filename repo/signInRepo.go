package repo

import (
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/utils/db"
	"log"
)

// SignInRepository implements SignInRepo interface for verifying user sign-in credentials.
type SignInRepository struct{}

// NewSignInRepository initializes and returns a new instance of SignInRepository.
func NewSignInRepository() *SignInRepository {
	return &SignInRepository{}
}

// VerifySignInCredentials verifies the user's sign-in credentials against the database.
func (repoStruct *SignInRepository) VerifySignInCredentials(username, password string) bool {
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
