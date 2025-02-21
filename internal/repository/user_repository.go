package repository

import (
	"gin-gorm-example/internal/model"

	"gorm.io/gorm"
)

// UserRepository 数据库访问层，封装用户相关的数据库操作
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository 创建新的 UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// GetUserByID 根据 ID 获取用户
func (r *UserRepository) GetUserByID(id string) (*model.User, error) {
	var user model.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
