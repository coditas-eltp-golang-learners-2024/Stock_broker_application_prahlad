package utils

import (
	"Stock_broker_application/src/app/authentication/models"
	"strings"

	"github.com/go-playground/validator/v10"
)

// SignUpValidations performs all the necessary validations for user signup.
func SignUpValidations(user *models.UserInfo) error {
	validate := validator.New()

	// Register custom validation functions if needed (already registered in ValidateCredentials)

	// Perform validation
	if err := validate.Struct(user); err != nil {
		return err // Validation error occurred
	}

	return nil // No validation error
}

// ValidateCredentials validates the provided SignInCredentials
func ValidateCredentials(creds models.SignInCredentials) error {
	validate := validator.New()

	// Register custom validation functions
	validate.RegisterValidation("containsDigit", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		for _, char := range password {
			if char >= '0' && char <= '9' {
				return true
			}
		}
		return false
	})

	validate.RegisterValidation("containsSpecial", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		specialChars := "!@#$%^&*()-_=+[]{};:'\"|\\<>,./?"
		for _, char := range password {
			if strings.ContainsRune(specialChars, char) {
				return true
			}
		}
		return false
	})

	// Perform validation
	err := validate.Struct(creds)
	if err != nil {
		return err // Validation error occurred
	}

	return nil // No validation error
}
