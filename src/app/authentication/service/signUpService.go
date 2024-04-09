package service

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/repo"
)

// UserService defines the interface for user-related operations.
type UserService interface {
	SignUp(user *models.UserInfo) error
}

type userService struct {
	userRepository repo.UserRepository
}

// NewUserService creates a new instance of UserService.
func NewUserService(userRepository repo.UserRepository) UserService {
	return &userService{userRepository}
}

// SignUp handles the user signup process.
func (s *userService) SignUp(user *models.UserInfo) error {
	// Perform custom validations
	if err := SignUpValidations(user); err != nil {
		return err
	}

	// Check if user already exists
	exists, err := repo.CheckUserExistenceByEmail(user.Email)
	if err != nil {
		return constants.ErrUserExistenceFailed
	}
	if exists {
		return constants.ErrUserAlreadyExists
	}

	// Insert user information into the database
	if err := s.userRepository.InsertUserInfo(*user); err != nil {
		return constants.ErrInsertUserInformation
	}

	return nil
}

// SignUpValidations performs all the necessary validations for user signup.
func SignUpValidations(user *models.UserInfo) error {
	// Your validation logic here
	// Example: validate name, email, phone number, etc.
	return nil
}
