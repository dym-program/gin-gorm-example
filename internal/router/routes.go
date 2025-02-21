package routes

import (
	"gin-gorm-example/internal/application"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置 Gin 路由
func SetupRoutes(router *gin.Engine) {
	router.GET("/image-path", application.GetImagePath) // 获取图片资源路径
	router.POST("/user", application.CreateUser)        // 创建用户
	router.GET("/users", application.GetAllUsers)       // 获取所有用户
	router.GET("/images", application.GetImageList)     // 获取所有图片
	router.POST("/upload", application.UploadImage)     // 上传图片
}
