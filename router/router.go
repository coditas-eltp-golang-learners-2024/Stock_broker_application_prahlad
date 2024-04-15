package router

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/docs"
	"Stock_broker_application/src/app/authentication/handlers"
	"Stock_broker_application/src/app/authentication/repo"
	"Stock_broker_application/src/app/authentication/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	signInRepo := repo.NewSignInRepository()
	signInService := service.NewSignInService(signInRepo)

	userRepository := repo.NewUserRepository()
	userService := service.NewUserService(userRepository)

	router.POST(constants.CreateUserRoute, handlers.PostUsersData(userService))
	router.POST(constants.SignInRoute, handlers.SignInHandler(signInService))

	// Adding Swagger documentation route
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
