package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController 控制器，处理 HTTP 请求
type UserController struct {
	UserService UserServiceInterface // 依赖 UserService 接口
}

// NewUserController 创建新的用户控制器
func NewUserController(userService UserServiceInterface) *UserController {
	return &UserController{UserService: userService}
}

// GetUser 获取用户信息
func (u *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := u.UserService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
