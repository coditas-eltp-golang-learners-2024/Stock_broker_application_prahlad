package signUpHandler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"stock_broker_application/constants"
	"stock_broker_application/utils/db/sqlSetUp"
)

var validate = validator.New()

func PostUsersData(c *gin.Context) {
	// Step 1: Binding the JSON format users request to readable format
	if err := c.ShouldBindJSON(&constants.UserInstance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Step 2: Checking if the user already exists through ID
	if exists := sqlSetUp.DB[constants.UserInstance.ID]; exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID already exists"})
		return
	}
	log.Println("New user identified!")

	// Step 3: Name validation of the given post body using validator
	if err := validate.Struct(constants.UserInstance); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			c.JSON(http.StatusBadRequest, gin.H{"error": e.Translate(validate)})
			return
		}
	}
	log.Println("User name validated!")

	// Step 4: Email validation of the post body using validator
	if err := validate.Var(constants.UserInstance.Email, "email"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}
	log.Println("Email validated!")
}
