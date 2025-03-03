package user

import (
	"fmt"
	"gin-gorm-example/internal/model"
	"gin-gorm-example/internal/repository"
)

// 定义接口，供控制器调用
type UserServiceInterface interface {
	GetUserByID(id string) (*model.User, error)
}

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) GetUserByID(id string) (*model.User, error) {
	user, err := s.UserRepo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("error fetching user: %v", err)
	}
	return user, nil
}
