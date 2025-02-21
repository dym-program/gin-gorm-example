package application

import (
	"fmt"
	"gin-gorm-example/configs"
	"gin-gorm-example/internal/model"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无法获取上传的文件"})
		return
	}

	uploadDir := configs.GlobalConfig.UploadDir
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		return
	}

	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	savePath := filepath.Join(uploadDir, filename)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the file"})
		return
	}

	imgURL := configs.GlobalConfig.ImagePath + filename
	image := model.Image{ImgName: filename, ImgURL: imgURL}
	if err := configs.DB.Create(&image).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store image info"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "图片上传成功", "image_url": imgURL})
}

func GetImageList(c *gin.Context) {
	var images []model.Image
	if err := configs.DB.Find(&images).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve images"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"images": images})
}
