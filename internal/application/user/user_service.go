package user

import (
	"fmt"
	"gin-gorm-example/internal/model"
	"gin-gorm-example/internal/repository"
)

// 定义服务接口，避免直接依赖具体实现
type UserServiceInterface interface {
	GetUserByID(id string) (*model.User, error)
}

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
