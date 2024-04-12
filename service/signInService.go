package service

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/repo"

	"Stock_broker_application/src/app/authentication/utils"
	"log"
)

// SignInService provides user sign-in operations using *repo.SignInRepository.
type SignInService struct {
	signInRepo *repo.SignInRepository
}

type SignInRepository *repo.SignInRepository

// NewSignInService creates a new instance of SignInService.
// @Summary Create a new instance of SignInService
// @Description Creates a new instance of SignInService with the provided signInRepo.
// @Accept json
// @Produce json
// @Param signInRepo body SignInRepository true "Pointer to SignInRepository"
// @Success 200 {string} string "Instance of SignInService created successfully"
// @Router /signin [post]
func NewSignInService(signInRepo SignInRepository) *SignInService {
	return &SignInService{signInRepo: signInRepo}
}

// VerifySignInCredentials verifies user's sign-in credentials.
// @Summary Verify user sign-in credentials
// @Description Verifies user's sign-in credentials.
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Param password path string true "Password"
// @Success 200 {string} string "true"
// @Failure 400 {string} string "false"
// @Router /verify [post]
func (s *SignInService) VerifySignInCredentials(username, password string) bool {
	return s.signInRepo.VerifySignInCredentials(username, password)
}

// SignInValidations performs validations on user sign-in credentials.
// @Summary Perform sign-in validations
// @Description Performs validations on user sign-in credentials and verifies them.
// @Accept json
// @Produce json
// @Param user body models.SignInCredentials true "User sign-in credentials"
// @Success 200 {string} string "Successfully signed in"
// @Failure 400 {string} string "Failed to sign in"
// @Router /signin [post]
func (s *SignInService) SignInValidations(user models.SignInCredentials) error {
	// Validate sign-in credentials
	err := utils.ValidateCredentials(user)
	if err != nil {
		log.Println("Error in validating user credentials:", err)
		return err // Return validation error
	}

	// Verify user credentials
	valid := s.signInRepo.VerifySignInCredentials(user.Name, user.Password)
	if !valid {
		log.Printf("Invalid username or password for user: %s\n", user.Name)
		return constants.ErrInVerification
	}

	return nil
}
