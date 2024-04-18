package constants

import "errors"

var (
	// User sign-up errors
	ErrInSignUp          = errors.New("failed to sign up")
	ErrUserAlreadyExists = errors.New("user with this email already exists")

	// User existence and data validation errors
	ErrUserExistenceFailed   = errors.New("failed to check user existence")
	ErrInvalidNameFormat     = errors.New("invalid name format, must be a string")
	ErrInvalidEmailFormat    = errors.New("invalid email format")
	ErrInvalidPhoneNumber    = errors.New("invalid phone number format or length (must be 10 numeric characters)")
	ErrInvalidPANCardFormat  = errors.New("invalid PAN card format or length (must be 10 alphanumeric characters)")
	ErrInvalidPasswordFormat = errors.New("invalid password format (must be at least 10 characters long and alphanumeric)")
	ErrInsertUserInformation = errors.New("failed to insert user information")
	ErrNoRowsAffected        = errors.New("no rows affected during user data insertion")

	// Server and internal errors
	ErrInternalError  = errors.New("internal server error")
	ErrStartingServer = errors.New("failed to start server")
	ErrInternalServer = errors.New("internal server error")

	// Authentication and verification errors
	ErrInVerification  = errors.New("invalid credentials, please try again")
	ErrIvalidOTPFormat = errors.New("invalid OTP format")

	// Binding and input data validation errors
	ErrInBindingData = errors.New("failed to bind data")
)
