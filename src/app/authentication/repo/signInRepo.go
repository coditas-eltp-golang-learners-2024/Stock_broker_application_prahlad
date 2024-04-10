package repo

import (
	"database/sql"
	"log"
)

type SignInRepo interface {
	VerifySignInCredentials(username, password string) bool
}

// Implementing Dependency Injection

// The below struct will help in creating an instance of db which we will recieve from the calling function.
type SignInRepository struct {
	db *sql.DB
}

// NewSignInRepository initializes and returns a new instance of SignInRepository with the provided *sql.DB instance.
// SignInRepository doesn't create its own database connection but relies on receiving one from the caller.
func NewSignInRepository(database *sql.DB) *SignInRepository {
	return &SignInRepository{db: database}
}

// Implementation of signInRepo interface
func (r *SignInRepository) VerifySignInCredentials(username, password string) bool {
	var userPasswordFromDB string
	err := r.db.QueryRow("SELECT password FROM userSignUpCredentials WHERE name = ?", username).Scan(&userPasswordFromDB)
	if err != nil {
		log.Println("Error verifying user credentials:", err)
		return false
	}

	// Compare hashed password with provided password using bcrypt
	if userPasswordFromDB == password {
		// Passwords match (authentication successful)
		return true
	} else {
		return false
	}
}
