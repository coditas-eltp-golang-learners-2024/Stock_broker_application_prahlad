package handlers

/*deal with HTTP request handling and response generation. They are responsible for processing incoming
requests, extracting necessary data, invoking appropriate services or business logic, and returning
 the response to the client.

*/
import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostUsersData(c *gin.Context) {
	// Step 1: Binding the JSON format users request to readable format
	if err := c.ShouldBindJSON(&constants.UserInstance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("Binding the data completed")

	//Step 2: Check if the users data already exist in the database
	exists, err := service.CheckIfUserExistsByEmail(constants.UserInstance.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking user existence"})
		return
	}

	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "User with this email already exists"})
		return
	}

	log.Println("New user verfied!")

	// Step 3: Validating name
	if !service.ValidateStringType(constants.UserInstance.Name) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid name format, must be a string"})
		return
	}
	log.Println("Name validation successful!")

	// Step 4: Validate the email format
	if !service.ValidateEmail(constants.UserInstance.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}
	log.Println("Email validation successful!")

	// Step 5: Validate the phone number length
	if !service.ValidatePANCard(constants.UserInstance.PanCardNumber) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid PAN card format or length (must be 10 alphanumeric characters)"})
		return
	}
	log.Println("Phone Number validation successful!")

	// Step 6: Validate PAN card format
	if !service.ValidatePANCard(constants.UserInstance.PanCardNumber) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid PAN card format or length (must be 10 alphanumeric characters)"})
		return
	}
	log.Println("PAN card number validation successful!")

	// Step 7: Validate password format and length
	if !service.ValidatePassword(constants.UserInstance.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 10 characters long and alphanumeric"})
		return
	}
	log.Println("Password validation successful!")

	// All validations passed, proceed to insert data into the database
	userInfo := models.UserInfo{
		ID:            constants.UserInstance.ID,
		Name:          constants.UserInstance.Name,
		Email:         constants.UserInstance.Email,
		PhoneNumber:   constants.UserInstance.PhoneNumber,
		PanCardNumber: constants.UserInstance.PanCardNumber,
		Password:      constants.UserInstance.Password,
	}

	// Insert data into the database
	err = service.InsertUserInfo(userInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting user data into the database"})
		return
	}

	log.Println("User data inserted successfully!")

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully!"})

}
