package main

import (
	"gin-gorm-example/configs"
	routes "gin-gorm-example/internal/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	configs.InitConfig()

	// 初始化数据库
	configs.InitDB()

	// 初始化 Gin 引擎
	r := gin.Default()

	// 配置静态文件目录（公开上传目录）
	r.Static(configs.GlobalConfig.ImagePath, configs.GlobalConfig.UploadDir) // 映射上传目录到静态 URL

	// 设置路由
	routes.SetupRoutes(r)

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
