package handlers

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/repo"
	"Stock_broker_application/src/app/authentication/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new user
// @Description Handles HTTP POST request to create new users
// @Accept json
// @Produce json
// @Param user body models.UserInfo true "User information in JSON format"
// @Success 200 {object} SuccessResponse "User information inserted successfully"
// @Failure 400 {object} ErrorResponse "Error message"
// @Router /signup [post]
func PostUsersData(c *gin.Context) {
	var user models.UserInfo

	// Bind request JSON to user model
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: constants.ErrInBindingData.Error()})
		return
	}

	// Create a new UserRepository instance
	userRepository := repo.NewUserRepository()

	// Create a new UserService instance with UserRepository injected
	userService := service.NewUserService(userRepository)

	// Call the SignUp method to handle user signup
	if err := userService.SignUp(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: constants.ErrInSignUp.Error()})
		return
	}

	// Respond with success message
	c.JSON(http.StatusOK, SuccessResponse{Message: "User information inserted successfully"})
}

// SuccessResponse represents a successful response format.
type SuccessResponses struct {
	Message string `json:"message"`
}

// ErrorResponse represents an error response format.
type ErrorResponses struct {
	Error string `json:"error"`
}
