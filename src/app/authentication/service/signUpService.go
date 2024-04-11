package service

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/repo"
	"Stock_broker_application/src/app/authentication/utils"
)

// UserService defines the interface for user-related operations.
type UserService interface {
	SignUp(user *models.UserInfo) error
}

type userService struct {
	userRepository repo.UserRepository
}

// NewUserService creates a new instance of UserService.
func NewUserService(UserRepository *repo.UserRepository) *userService {
	return &userService{userRepository: *UserRepository}
}

// SignUp handles the user signup process.
func (s *userService) SignUp(user *models.UserInfo) error {
	// Perform custom validations
	if err := utils.SignUpValidations(user); err != nil {
		return err
	}

	// Check if user already exists
	exists, err := s.userRepository.CheckIfUserExistsByEmail(user.Email)
	if err != nil {
		return constants.ErrUserExistenceFailed
	}
	if exists {
		return constants.ErrUserAlreadyExists
	}

	// Insert user information into the database
	if err := s.userRepository.InsertUserInfo(user); err != nil {
		return constants.ErrInsertUserInformation
	}

	return nil
}

// SignUpValidations performs all the necessary validations for user signup.
func SignUpValidations(user *models.UserInfo) error {

	return nil
}
