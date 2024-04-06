package service

import (
	"database/sql"
	"log"
	"reflect"
	"regexp"

	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/utils"

	"github.com/go-playground/validator/v10"
)

var (
	emailRegex        = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	panRegex          = regexp.MustCompile(`^[A-Za-z]{5}[0-9]{4}[A-Za-z]{1}$`)
	alphanumericRegex = regexp.MustCompile("^[a-zA-Z0-9]+$")
)

/*They implement the actual business rules, perform data manipulation,
interact with databases or external services, etc.
 It abstracts away the details of data access, validation, and other low-level concerns.
*/

// Step 2: Check if the provided email already exists in the map
func CheckIfUserExistsByEmail(email string) (bool, error) {
	// Prepare SQL query to check if the email exists
	query := "SELECT COUNT(*) FROM userSignUpCredentials WHERE email = ?"

	// Execute the query
	var count int
	err := utils.Db.QueryRow(query, email).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			// No user with this email exists
			return false, nil
		}
		log.Printf("Error checking user existence: %v", err)
		return false, err
	}

	// If count > 0, user with this email exists
	return count > 0, nil
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
func InsertUserInfo(userInfo models.UserInfo) error {
	// Prepare SQL query for insertion
	query := `
		INSERT INTO userSignUpCredentials ( id, name, email, phoneNumber, panCardNumber, password)
		VALUES (?,?, ?, ?, ?, ?)
	`

	// Execute SQL query with prepared statement
	result, err := utils.Db.Exec(query, userInfo.ID, userInfo.Name, userInfo.Email, userInfo.PhoneNumber, userInfo.PanCardNumber, userInfo.Password)
	if err != nil {
		log.Printf("Error inserting user data into the database: %v", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		log.Println("No rows affected during user data insertion")
		return sql.ErrNoRows
	}

	return nil
}
