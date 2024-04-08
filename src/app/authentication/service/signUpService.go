package service

import (
	"reflect"
	"regexp"

	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/repo"

	"github.com/go-playground/validator/v10"
)

var (
	emailRegex        = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	panRegex          = regexp.MustCompile(`^[A-Za-z]{5}[0-9]{4}[A-Za-z]{1}$`)
	alphanumericRegex = regexp.MustCompile("^[a-zA-Z0-9]+$")
)

// UserService handles user-related business logic.
type UserService struct {
}

// NewUserService creates a new instance of UserService.
func NewUserService() *UserService {
	return &UserService{}
}

/*They implement the actual business rules, perform data manipulation,
interact with databases or external services, etc.
 It abstracts away the details of data access, validation, and other low-level concerns.
*/

// Step 2: Check if the provided email already exists in the map
func (s *UserService) CheckUserExistenceByEmail(email string) (bool, error) {
	// Call the data access layer (repository) function to check user existence
	return repo.CheckIfUserExistsByEmail(email)
}

// ValidateStringType checks if the input parameter is of type string
func ValidateStringType(input interface{}) bool {
	// Use reflection to determine the type of the input
	return reflect.TypeOf(input).Kind() == reflect.String
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateEmail checks if the provided email is in a valid format
func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func ValidatePhoneNumber(phoneNumber string) bool {
	// Check if the phone number is exactly 10 digits and consists only of numeric characters
	if len(phoneNumber) != 10 {
		return false
	}

	for _, char := range phoneNumber {
		if char < '0' || char > '9' {
			return false // Non-numeric character found
		}
	}

	return true
}

// ValidatePANCard checks if the provided PAN card number is exactly 10 characters long and alphanumeric
func ValidatePANCard(panCardNumber string) bool {
	// Check PAN card number length
	if len(panCardNumber) != 10 {
		return false
	}

	// Check if PAN card number is alphanumeric
	return alphanumericRegex.MatchString(panCardNumber)
}

// ValidatePassword checks if the provided password is at least 10 characters long and alphanumeric
func ValidatePassword(password string) bool {
	// Check password length
	if len(password) < 10 {
		return false
	}

	// Check if password is alphanumeric
	return alphanumericRegex.MatchString(password)
}

// InsertUserInfo inserts user information into the database
// InsertUser inserts user information into the database.
func (s *UserService) InsertUser(userInfo models.UserInfo) error {
	// Call the data access layer (repository) function to insert user information
	return repo.InsertUserInfo(userInfo)
}
