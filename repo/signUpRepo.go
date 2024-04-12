package repo

import (
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/utils/db"
)

// UserRepository represents a repository for user-related database operations.
type UserRepository struct{}

// NewUserRepository creates a new instance of UserRepository.
// @return *UserRepository
func NewUserRepository() *UserRepository {
	return &UserRepository{} // Return a pointer to a new instance of UserRepository
}

// CheckIfUserExistsByEmail checks if a user with the given email exists in the database.
// @param email path string true "User email"
// @return bool
// @return error
func (r *UserRepository) CheckIfUserExistsByEmail(email string) (bool, error) {
	var count int64
	result := db.GormDb.Model(&models.UserInfo{}).Where("email = ?", email).Count(&count)
	if result.Error != nil {
		return false, result.Error // Return false and error if query execution fails
	}
	return count > 0, nil // Return true if user exists, false otherwise
}

// InsertUserInfo inserts user information into the database.
// @param userInfo body models.UserInfo true "User information"
// @return error
func (r *UserRepository) InsertUserInfo(userInfo *models.UserInfo) error {
	result := db.GormDb.Create(userInfo)
	if result.Error != nil {
		return result.Error // Return error if insertion fails
	}
	if result.RowsAffected == 0 {
		return db.GormDb.Error // Return database error if no rows were affected
	}
	return nil // Return nil (no error) on successful insertion
}

// CheckUserExistenceByEmail checks if a user with the given email exists in the database.
// @param email path string true "User email"
// @return bool
// @return error
func CheckUserExistenceByEmail(email string) (bool, error) {
	userRepository := NewUserRepository()
	return userRepository.CheckIfUserExistsByEmail(email)
}
