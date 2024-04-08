package repo

import (
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/utils" // Import the database connection utility
	"database/sql"
	"log"
)

// CheckIfUserExistsByEmail checks if a user with the given email exists in the database.
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

// InsertUserInfo inserts user information into the database.
func InsertUserInfo(userInfo models.UserInfo) error {
	// Prepare SQL query for insertion
	query := `
		INSERT INTO userSignUpCredentials (id, name, email, phoneNumber, panCardNumber, password)
		VALUES (?, ?, ?, ?, ?, ?)
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
