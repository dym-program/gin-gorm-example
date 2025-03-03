package routes

import (
	"gin-gorm-example/internal/application/image"

	"github.com/gin-gonic/gin"
)

func SetupImageRoutes(router *gin.Engine, imageController *image.ImageController) {
	router.GET("/images", imageController.GetImageList)
	router.POST("/upload", imageController.UploadImage)
}
