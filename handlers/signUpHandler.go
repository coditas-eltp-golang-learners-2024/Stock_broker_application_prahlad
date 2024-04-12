package handlers

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostUsersData is a handler function for creating a new user.
// It accepts a UserService instance and Gin context, then handles the POST request.
// @Summary Create a new user
// @Description Handles HTTP POST request to create new users
// @Accept json
// @Produce json
// @Param user body models.UserInfo true "User information in JSON format"
// @Success 200 {object} models.SuccessResponse "User information inserted successfully"
// @Failure 400 {object} models.ErrorResponse "Error message"
// @Router /signup [post]
func PostUsersData(userService *service.UserServiceStruct) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.UserInfo

		// Bind request JSON to user model
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: constants.ErrInBindingData.Error()})
			return
		}

		// Call the SignUp method to handle user signup
		if err := userService.SignUp(&user); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: constants.ErrInSignUp.Error()})
			return
		}

		// Respond with success message
		c.JSON(http.StatusOK, models.SuccessResponse{Message: "User information inserted successfully"})
	}
}
