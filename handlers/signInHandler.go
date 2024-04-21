package handlers

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/repo"
	"Stock_broker_application/src/app/authentication/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary User sign-in
// @Description Handles user sign-in endpoint
// @Accept json
// @Produce json
// @Param credentials body models.SignInCredentials true "User credentials for sign-in"
// @Success 200 {object} models.SuccessResponse "Authentication successful"
// @Failure 400 {object} models.ErrorResponse "Invalid credentials or error message"
// @Router /signin [post]

func SignInHandler(signInService *service.SignInService, otpService service.OTPValidationService, otpRepo repo.OTPValidationRepo) gin.HandlerFunc {
	return func(context *gin.Context) {
		var userCredentials models.SignInCredentials

		// Binding the body of the incoming POST request
		if err := context.ShouldBindJSON(&userCredentials); err != nil {
			context.JSON(http.StatusBadRequest, models.ErrorResponse{Error: constants.ErrInBindingData.Error()})
			return
		}

		// Call the SignInValidations method to validate user credentials
		if err := signInService.SignInValidations(userCredentials); err != nil {
			context.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid credentials, please try again"})
			return
		}

		context.JSON(http.StatusOK, models.SuccessResponse{Message: "Authentication successful"})

		//Storing the genearted otp in database after successful SignIn
		// Generate OTP using OTP service
		otp, createdAt := otpService.GenerateOTP()
		otpInstance := models.OTP{
			Email: userCredentials.Email,
		}

		// Store the generated OTP in the database
		if err := otpRepo.InsertOTP(otp, createdAt, otpInstance.Email); err != nil {
			context.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to store OTP, please try again"})
			return
		}

	}
}
