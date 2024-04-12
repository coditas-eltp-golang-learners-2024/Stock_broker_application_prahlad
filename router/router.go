package router

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/handlers"
	"Stock_broker_application/src/app/authentication/repo"
	"Stock_broker_application/src/app/authentication/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetUpRouter sets up the Gin router with API endpoints and Swagger documentation.
// @title Stock Broker Application
// @version 1.0
// @description This is a Stock Broker Application API
// @host localhost:8080
// @BasePath /
func SetUpRouter() *gin.Engine {
	router := gin.Default()

	// Define route for handling POST requests to create new users
	// @Summary Create a new user
	// @Description Handles HTTP POST request to create new users
	// @Accept json
	// @Produce json
	// @Router /users [post]
	router.POST(constants.CreateUserRoute, handlers.PostUsersData)

	// Create SignInRepository instance with database connection
	signInRepo := repo.NewSignInRepository()

	// Create SignInService instance with SignInRepository dependency injected
	signInService := service.NewSignInService(signInRepo)

	// Define route for handling POST requests to sign in users
	// @Summary User sign-in
	// @Description Handles user sign-in endpoint
	// @Accept json
	// @Produce json
	// @Param credentials body models.SignInCredentials true "User credentials for sign-in"
	// @Router /signin [post]
	router.POST(constants.SignInRoute, handlers.SignInHandler(signInService))

	// Adding Swagger documentation route
	// @Router /docs/*any [get]
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
