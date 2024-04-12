package constants

import "errors"

// Custom error types used throughout the application
var (
	// ErrUserExistenceFailed indicates a failure in checking user existence.
	ErrUserExistenceFailed = errors.New("failed to check user existence")

	// ErrInvalidNameFormat indicates an invalid name format (must be a string).
	ErrInvalidNameFormat = errors.New("invalid name format, must be a string")

	// ErrInvalidEmailFormat indicates an invalid email format.
	ErrInvalidEmailFormat = errors.New("invalid email format")

	// ErrInvalidPhoneNumber indicates an invalid phone number format or length (must be 10 numeric characters).
	ErrInvalidPhoneNumber = errors.New("invalid phone number format or length (must be 10 numeric characters)")

	// ErrInvalidPANCardFormat indicates an invalid PAN card format or length (must be 10 alphanumeric characters).
	ErrInvalidPANCardFormat = errors.New("invalid PAN card format or length (must be 10 alphanumeric characters)")

	// ErrInvalidPasswordFormat indicates an invalid password format (must be at least 10 characters long and alphanumeric).
	ErrInvalidPasswordFormat = errors.New("invalid password format (must be at least 10 characters long and alphanumeric)")

	// ErrInsertUserInformation indicates a failure in inserting user information.
	ErrInsertUserInformation = errors.New("failed to insert user information")

	// ErrNoRowsAffected indicates no rows were affected during user data insertion.
	ErrNoRowsAffected = errors.New("no rows affected during user data insertion")

	// ErrUserAlreadyExists indicates that a user with the specified email already exists.
	ErrUserAlreadyExists = errors.New("user with this email already exists")

	// ErrInternalError indicates an internal server error.
	ErrInternalError = errors.New("internal server error")

	// ErrStartingServer indicates a failure to start the server.
	ErrStartingServer = errors.New("failed to start server")

	// ErrInVerification indicates invalid credentials during verification.
	ErrInVerification = errors.New("invalid credentials, please try again")

	// ErrInBindingData indicates a failure to bind data.
	ErrInBindingData = errors.New("failed to bind data")

	// ErrInSignUp indicates a failure in signing up.
	ErrInSignUp = errors.New("Failed to sign up")
)
