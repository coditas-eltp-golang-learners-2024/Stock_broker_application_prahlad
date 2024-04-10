package router

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/handlers"
	"Stock_broker_application/src/app/authentication/repo"
	"Stock_broker_application/src/app/authentication/service"
	dbpackage "Stock_broker_application/src/app/authentication/utils/db"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	// Define route for handling POST requests to create new users
	router.POST(constants.CreateUserRoute, handlers.PostUsersData)

	// Create SignInRepository instance with database connection
	signInRepo := repo.NewSignInRepository(dbpackage.Db)

	// Create SignInService instance with SignInRepository dependency injected
	signInService := service.NewUserSignInService(signInRepo)

	// Define route for handling POST requests to sign in users
	router.POST(constants.SignInRoute, handlers.SignInHandler(signInService))

	return router
}
