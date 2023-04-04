package main

import (
	"AppServer/config"
	Todo "AppServer/services/todo"

	"context"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	engineConfig := config.InitAppConfig()
	mongoClient, mongoClientError := mongo.Connect(context.TODO(), options.Client().ApplyURI(
		fmt.Sprintf("%s/%s", engineConfig.MONGO_URI, engineConfig.DATABASE),
	))
	fmt.Printf("%s/%s", engineConfig.MONGO_URI, engineConfig.DATABASE)
	if mongoClientError != nil {
		panic(mongoClientError)
	}
	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	gin.SetMode(engineConfig.GIN_MODE)
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	Todo.TodoController{
		TodoService: &Todo.TodoService{
			MongoContext: mongoClient,
			DATABASE:     engineConfig.DATABASE,
			COLLECTION:   engineConfig.COLLECTION,
		},
	}.RegisterController(engine)

	fmt.Println(engine.Routes())
	engine.Run(fmt.Sprintf(":%d", engineConfig.PORT))

}
