package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService UserServiceInterface
}

func NewUserController(service UserServiceInterface) *UserController {
	return &UserController{UserService: service}
}

func (uc *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := uc.UserService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
