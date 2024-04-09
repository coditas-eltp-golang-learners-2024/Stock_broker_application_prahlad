package repo

import (
	"Stock_broker_application/src/app/authentication/models"
	utils "Stock_broker_application/src/app/authentication/utils/db"
	"database/sql"
	"log"
)

type UserRepository struct{}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository() *UserRepository {
	return &UserRepository{} // Return a pointer to a new instance of UserRepository
}

func (r *UserRepository) CheckIfUserExistsByEmail(email string) (bool, error) {
	query := "SELECT COUNT(*) FROM userSignUpCredentials WHERE email = ?"
	var count int
	err := utils.Db.QueryRow(query, email).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		log.Printf("Error checking user existence: %v", err)
		return false, err
	}
	return count > 0, nil
}

func (r *UserRepository) InsertUserInfo(userInfo models.UserInfo) error {
	query := `
        INSERT INTO userSignUpCredentials (id, name, email, phoneNumber, panCardNumber, password)
        VALUES (?, ?, ?, ?, ?, ?)
    `
	result, err := utils.Db.Exec(query, userInfo.ID, userInfo.Name, userInfo.Email, userInfo.PhoneNumber, userInfo.PanCardNumber, userInfo.Password)
	if err != nil {
		log.Printf("Error inserting user data into the database: %v", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error retrieving rows affected:", err)
		return err
	}
	if rowsAffected == 0 {
		log.Println("No rows affected during user data insertion")
		return sql.ErrNoRows
	}
	return nil
}

// CheckUserExistenceByEmail checks if a user with the given email exists.
func CheckUserExistenceByEmail(email string) (bool, error) {
	userRepository := NewUserRepository()
	return userRepository.CheckIfUserExistsByEmail(email)
}
