package handlers

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary User sign-in
// @Description Handles user sign-in endpoint
// @Accept json
// @Produce json
// @Param credentials body models.SignInCredentials true "User credentials for sign-in"
// @Success 200 {object} SuccessResponse "Authentication successful"
// @Failure 400 {object} ErrorResponse "Invalid credentials or error message"
// @Router /signin [post]
func SignInHandler(signInService *service.SignInService) gin.HandlerFunc {
	return func(context *gin.Context) {
		var userCredentials models.SignInCredentials

		// Binding the body of the incoming POST request
		if err := context.ShouldBindJSON(&userCredentials); err != nil {
			context.JSON(http.StatusBadRequest, ErrorResponse{Error: constants.ErrInBindingData.Error()})
			return
		}

		// Call the SignInValidations method to validate user credentials
		if err := signInService.SignInValidations(userCredentials); err != nil {
			context.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid credentials, please try again"})
			return
		}

		context.JSON(http.StatusOK, SuccessResponse{Message: "Authentication successful"})
	}
}

// SuccessResponse represents a successful response format.
type SuccessResponse struct {
	Message string `json:"message"`
}

// ErrorResponse represents an error response format.
type ErrorResponse struct {
	Error string `json:"error"`
}
