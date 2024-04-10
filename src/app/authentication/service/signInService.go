package service

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/repo"
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

func (u *UserSignInService) SignInValidations(user models.SignInCredentials) error {
	valid := u.signInRepo.VerifySignInCredentials(user.UserName, user.Password)

	if !valid {
		log.Printf("Invalid username or password for user: %s\n", user.UserName)

		return constants.ErrInVerification
	}

	return nil
}
