// handlers/post_users_data.go

package handlers

import (
	"log"
	"net/http"

	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/service"

	"github.com/gin-gonic/gin"
)

func PostUsersData(c *gin.Context) {
	var users []models.UserInfo

	// Step 1: Binding the JSON format users request to a readable format
	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Binding of data completed. Number of users: %d", len(users))

	// Create channels for communication and synchronization
	errorChan := make(chan error, len(users))
	doneChan := make(chan struct{})

	// Process each user signup concurrently using goroutines
	for _, user := range users {
		go func(u models.UserInfo) {
			defer func() {
				if r := recover(); r != nil {
					errorChan <- constants.ErrInternalError
				}
			}()

			// Step 2: Validate and process each user signup
			err := processUserSignup(u)
			if err != nil {
				errorChan <- err
			}
		}(user)
	}

	// Use a goroutine to wait for all user signups to complete
	go func() {
		for range users {
			select {
			case err := <-errorChan:
				if err != nil {
					doneChan <- struct{}{} // Signal completion with error
					return
				}
			}
		}
		doneChan <- struct{}{} // Signal successful completion
	}()

	// Wait for all user signups to complete or encounter an error
	select {
	case <-doneChan:
		c.JSON(http.StatusOK, gin.H{"message": "All user information inserted successfully"})
		log.Println("All user information inserted successfully")
	case <-c.Request.Context().Done():
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Request cancelled or timed out"})
		log.Println("Request cancelled or timed out")
	}

	// Close channels to prevent goroutine leaks
	close(errorChan)
	close(doneChan)
}

func processUserSignup(user models.UserInfo) error {
	// Step 3: Validate user data
	if !service.ValidateStringType(user.Name) {
		return constants.ErrInvalidNameFormat
	}
	if !service.ValidateEmail(user.Email) {
		return constants.ErrInvalidEmailFormat
	}
	if !service.ValidatePhoneNumber(user.PhoneNumber) {
		return constants.ErrInvalidPhoneNumber
	}
	if !service.ValidatePANCard(user.PanCardNumber) {
		return constants.ErrInvalidPANCardFormat
	}
	if !service.ValidatePassword(user.Password) {
		return constants.ErrInvalidPasswordFormat
	}

	// Step 4: Insert user information into the database
	userService := service.NewUserService()
	if err := userService.InsertUser(user); err != nil {
		return constants.ErrInsertUserInformation
	}

	return nil
}
