package Todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoService struct {
	MongoContext *mongo.Client
	DATABASE     string
	COLLECTION   string
}

// Create a new Task
func (service TodoService) CreateOne(context *gin.Context) {

	var task Todo

	if err := context.BindJSON(&task); err != nil {
		return
	}
	task.InsertOne(service.MongoContext, service.DATABASE, service.COLLECTION)
	context.IndentedJSON(http.StatusCreated, task)
}

// Find all Created Taskes
func (service TodoService) FindAll(context *gin.Context) {

	context.IndentedJSON(http.StatusCreated, Todo{}.FindAll(service.MongoContext, service.DATABASE, service.COLLECTION))
}

// Find a task by a given Id
func (service TodoService) FindOne(context *gin.Context) {

	var task Todo

	if err := context.BindJSON(&task); err != nil {
		return
	}
	task.FindOne(service.MongoContext, service.DATABASE, service.COLLECTION)

	context.IndentedJSON(http.StatusCreated, task)
}

// Update a Task by a given Id
func (service TodoService) UpdateOne(context *gin.Context) {

	var task Todo
	if err := context.BindJSON(&task); err != nil {
		return
	}
	task.UpdateOne(service.MongoContext, service.DATABASE, service.COLLECTION)

	context.IndentedJSON(http.StatusCreated, task)
}

// Delete a Task by a given Id
func (service TodoService) DeleteOne(context *gin.Context) {

	var task Todo
	if err := context.BindJSON(&task); err != nil {
		return
	}
	task.DeleteOne(service.MongoContext, service.DATABASE, service.COLLECTION)

	context.IndentedJSON(http.StatusCreated, task)
}
