package service

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/repo"
	"Stock_broker_application/src/app/authentication/utils"
	"log"
)

type UserSignInService struct {
	signInRepo repo.SignInRepo
}

func NewUserSignInService(signInRepo repo.SignInRepo) *UserSignInService {
	return &UserSignInService{
		signInRepo: signInRepo,
	}
}

// below function first validates and then verfies the user credentials
func (u *UserSignInService) SignInValidations(user models.SignInCredentials) error {

	//Firstly Validating signIn Credentials
	err := utils.ValidateCredentials(user)
	if err != nil {
		log.Println("Error in validating user credentials:", err)
		return err // Return validation error
	}

	// Call the VerifySignInCredentials method to validate user credentials
	valid := u.signInRepo.VerifySignInCredentials(user.Name, user.Password)

	if !valid {
		log.Printf("Invalid username or password for user: %s\n", user.Name)

		return constants.ErrInVerification
	}

	return nil
}
