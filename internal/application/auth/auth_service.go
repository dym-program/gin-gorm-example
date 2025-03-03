package auth

import (
	"fmt"
	"gin-gorm-example/internal/model"
	"gin-gorm-example/internal/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	Register(username, email, password string) error
	Login(username, password string) (string, error)
}

type AuthService struct {
	UserRepo *repository.UserRepository
	jwtKey   []byte
}

func NewAuthService(userRepo *repository.UserRepository, jwtKey []byte) *AuthService {
	return &AuthService{
		UserRepo: userRepo,
		jwtKey:   jwtKey,
	}
}

func (s *AuthService) Register(username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	user := &model.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}
	return s.UserRepo.Create(user)
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.UserRepo.FindByUsername(username)
	if err != nil {
		return "", fmt.Errorf("user not found: %v", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("invalid password: %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(72 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(s.jwtKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return tokenString, nil
}
