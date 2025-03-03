package routes

import (
	"gin-gorm-example/internal/application/user"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, userController *user.UserController) {
	router.GET("/user/:id", userController.GetUser)
}
