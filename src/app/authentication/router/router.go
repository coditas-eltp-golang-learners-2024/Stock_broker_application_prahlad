package router

import (
	"Stock_broker_application/src/app/authentication/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	// Define route for handling POST requests to create new users
	router.POST("/users", handlers.PostUsersData)

	return router
}