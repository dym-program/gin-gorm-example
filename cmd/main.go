package main

import (
	"gin-gorm-example/configs"
	"gin-gorm-example/internal/application/user"
	"gin-gorm-example/internal/repository"
	routes "gin-gorm-example/internal/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置和数据库
	configs.InitConfig()
	err := configs.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 初始化数据库访问层、服务层和控制器
	userRepo := repository.NewUserRepository(configs.DB)
	userService := user.NewUserService(userRepo)
	userController := user.NewUserController(userService)

	// 初始化 Gin 引擎
	r := gin.Default()

	// 设置路由
	routes.SetupRoutes(r, userController)

	// 启动服务器
	r.Run(":8080")
}
