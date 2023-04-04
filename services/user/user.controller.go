package user

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *UserService
}

func (userController UserController) RegisterController(engine *gin.Engine) {

	userRegistry := engine.Group("/user")
	{
		userRegistry.POST("/create", userController.UserService.SaveUser)
	}

}
