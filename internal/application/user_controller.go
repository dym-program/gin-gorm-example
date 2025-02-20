package application

import (
	"gin-gorm-example/configs"
	"gin-gorm-example/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetImagePath 返回图片资源地址
func GetImagePath(c *gin.Context) {
	imagePath := configs.GlobalConfig.ImagePath
	c.JSON(http.StatusOK, gin.H{"image_path": imagePath})
}

// CreateUser 创建用户
func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 保存用户到数据库
	if err := configs.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user created", "user": user})
}

// GetAllUsers 获取所有用户
func GetAllUsers(c *gin.Context) {
	var users []model.User
	if err := configs.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
