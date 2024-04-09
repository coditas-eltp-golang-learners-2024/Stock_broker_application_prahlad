package constants

import "errors"

// Custom error types
var (
	ErrUserExistenceFailed   = errors.New("failed to check user existence")
	ErrInvalidNameFormat     = errors.New("invalid name format, must be a string")
	ErrInvalidEmailFormat    = errors.New("invalid email format")
	ErrInvalidPhoneNumber    = errors.New("invalid phone number format or length (must be 10 numeric characters)")
	ErrInvalidPANCardFormat  = errors.New("invalid PAN card format or length (must be 10 alphanumeric characters)")
	ErrInvalidPasswordFormat = errors.New("invalid password format (must be at least 10 characters long and alphanumeric)")
	ErrInsertUserInformation = errors.New("failed to insert user information")
	ErrNoRowsAffected        = errors.New("no rows affected during user data insertion")
	ErrUserAlreadyExists     = errors.New("user with this email already exists")
	ErrInternalError         = errors.New("internal server error")
	ErrStartingServer        = errors.New("failed to start server")
)
