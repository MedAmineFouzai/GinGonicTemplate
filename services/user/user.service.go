package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	MongoContext *mongo.Client
}

func (userService UserService) SaveUser(context *gin.Context) {

	var user User

	if err := context.BindJSON(&user); err != nil {
		return
	}
	user.Save(userService.MongoContext)
	context.IndentedJSON(http.StatusCreated, user)
}
