package routes

import (
	"gin-gorm-example/internal/application/auth"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine, authController *auth.AuthController) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", authController.RegisterHandler)
		authGroup.POST("/login", authController.LoginHandler)
		// 可扩展：找回密码、重置密码等
	}
}
