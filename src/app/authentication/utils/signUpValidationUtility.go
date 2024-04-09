package utils

import (
	//"Stock_broker_application/src/app/authentication/repo"
	//"Stock_broker_application/src/app/authentication/repo"
	"reflect"
	"regexp"
)

// ValidateStringType checks if the input parameter is of type string.
func ValidateStringType(input interface{}) bool {
	return reflect.TypeOf(input).Kind() == reflect.String
}

// ValidateEmail checks if the provided email is in a valid format.
func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// ValidatePhoneNumber checks if the provided phone number is in a valid format.
func ValidatePhoneNumber(phoneNumber string) bool {
	if len(phoneNumber) != 10 {
		return false
	}
	return true
}

// ValidatePANCard checks if the provided PAN card number is in a valid format.
func ValidatePANCard(panCardNumber string) bool {
	if len(panCardNumber) != 10 {
		return false
	}
	return true
}

// ValidatePassword checks if the provided password is in a valid format.
func ValidatePassword(password string) bool {
	if len(password) < 10 {
		return false
	}
	return true
}
