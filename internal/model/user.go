package model

import "gorm.io/gorm"

// User 模型定义
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}
