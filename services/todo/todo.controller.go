package Todo

import (
	"github.com/gin-gonic/gin"
)

type TodoController struct {
	TodoService *TodoService
}

func (controller TodoController) RegisterController(engine *gin.Engine) {

	TodoRegistry := engine.Group("/Todo")
	{
		TodoRegistry.POST("/CreateOne", controller.TodoService.CreateOne)
		TodoRegistry.GET("/FindAll", controller.TodoService.FindAll)
		TodoRegistry.GET("/FindOne", controller.TodoService.FindOne)
		TodoRegistry.PUT("/UpdateOne", controller.TodoService.UpdateOne)
		TodoRegistry.DELETE("/DeleteOne", controller.TodoService.DeleteOne)

	}

}
