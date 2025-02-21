package routes

import (
	"gin-gorm-example/internal/application/user"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置 Gin 路由
func SetupRoutes(router *gin.Engine, userController *user.UserController) {
	router.GET("/user/:id", userController.GetUser)
}
