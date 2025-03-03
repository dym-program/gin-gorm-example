package main

import (
	"gin-gorm-example/configs"
	"gin-gorm-example/internal/application/auth"
	imageApp "gin-gorm-example/internal/application/image"
	userApp "gin-gorm-example/internal/application/user"
	"gin-gorm-example/internal/repository"
	routes "gin-gorm-example/internal/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置和数据库
	configs.InitConfig()
	if err := configs.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 初始化 Repositories
	userRepo := repository.NewUserRepository(configs.DB)
	imageRepo := repository.NewImageRepository(configs.DB)

	// JWT 密钥（应从配置中读取，这里简单写死）
	jwtKey := []byte("your_secret_key")

	// 初始化 Auth 模块：服务和控制器
	authService := auth.NewAuthService(userRepo, jwtKey)
	authController := auth.NewAuthController(authService)

	// 初始化 User 模块：服务和控制器
	userService := userApp.NewUserService(userRepo)
	userController := userApp.NewUserController(userService)

	// 初始化 Image 模块：服务和控制器
	imageService := imageApp.NewImageService(imageRepo, configs.GlobalConfig.UploadDir, configs.GlobalConfig.ImagePath)
	imageController := imageApp.NewImageController(imageService)

	// 初始化 Gin 引擎
	r := gin.Default()

	// 配置静态文件服务：将上传目录映射到 /static/images 供前端下载图片
	r.Static("/static/images", configs.GlobalConfig.UploadDir)

	// 注册各个模块的路由
	routes.SetupAuthRoutes(r, authController)
	routes.SetupUserRoutes(r, userController)
	routes.SetupImageRoutes(r, imageController)

	// 启动服务器
	r.Run(":8080")
}
