package main

import (
	"gin-gorm-example/configs"
	"gin-gorm-example/internal/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// 初始化配置
	configs.InitConfig()

	// 初始化数据库
	configs.InitDB()

	// 初始化 Gin 引擎
	r := gin.Default()

	// 设置路由
	routes.SetupRoutes(r)

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
