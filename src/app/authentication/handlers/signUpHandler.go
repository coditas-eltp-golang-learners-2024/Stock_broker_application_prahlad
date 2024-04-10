package handlers

import (
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/repo"
	"Stock_broker_application/src/app/authentication/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostUsersData handles the HTTP POST request to create new users.
func PostUsersData(c *gin.Context) {
	var user models.UserInfo

	// Bind request JSON to user model
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new UserRepository instance
	userRepository := repo.NewUserRepository()

	// Create a new UserService instance with UserRepository injected
	userService := service.NewUserService(userRepository)

	// Call the SignUp method to handle user signup
	if err := userService.SignUp(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "User information inserted successfully"})
}
