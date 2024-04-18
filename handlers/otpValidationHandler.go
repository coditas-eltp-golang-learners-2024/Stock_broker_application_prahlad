package handlers

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/repo"
	"Stock_broker_application/src/app/authentication/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ValidateOTPHandler is a handler function for validating OTP.
// @Summary Validate OTP
// @Description Validate OTP using the provided service and repository
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body models.ValidateOTPRequest true "OTP validation request"
// @Param otpService query service.OTPValidationService true "OTP validation service"
// @Router /validateOtp [post]
func ValidateOTPHandler(otpService service.OTPValidationService, otpRepo repo.OTPValidationRepo) gin.HandlerFunc {
	return func(context *gin.Context) {
		var req models.ValidateOTPRequest
		if err := context.ShouldBindJSON(&req); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate OTP using the provided OTPValidationService and OTPValidationRepo
		valid, err := otpService.ValidateOTP(req.OTP, req.Email, otpRepo)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": constants.ErrInternalServer})
			return
		}

		// Respond with OTP validation results
		if valid {
			context.JSON(http.StatusOK, gin.H{"message": "OTP is valid"})
		} else {
			context.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrIvalidOTP})
		}
	}
}
