package handlers

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignInHandler is a handler function for user sign-in endpoint.
func SignInHandler(signInService *service.UserSignInService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userCredentials models.SignInCredentials

		// Binding the body of the incoming POST request
		if err := c.ShouldBindJSON(&userCredentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInBindingData})
			return
		}

		// Call the SignInValidations method to validate user credentials
		if err := signInService.SignInValidations(userCredentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials, please try again"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Authentication successful"})
	}
}
