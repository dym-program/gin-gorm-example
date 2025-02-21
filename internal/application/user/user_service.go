package application

import (
	"fmt"
	"gin-gorm-example/internal/model"
	"gin-gorm-example/internal/repository"
)

// UserService 服务层，封装用户相关的业务逻辑
type UserService struct {
	UserRepo *repository.UserRepository
}

// NewUserService 创建新的 UserService
func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

// GetUserByID 根据 ID 获取用户
func (s *UserService) GetUserByID(id string) (*model.User, error) {
	user, err := s.UserRepo.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("error fetching user: %v", err)
	}
	return user, nil
}
