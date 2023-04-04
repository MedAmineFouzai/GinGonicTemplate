package main

import (
	"AppServer/config"
	"AppServer/services/user"
	"context"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	engineConfig := config.InitAppConfig()
	mongoClient, mongoClientError := mongo.Connect(context.Background(), options.Client().ApplyURI(
		fmt.Sprintf("%s/%s", engineConfig.MONGO_URI, engineConfig.DATABASE),
	))
	fmt.Printf("%s/%s", engineConfig.MONGO_URI, engineConfig.DATABASE)
	if mongoClientError != nil {
		panic(mongoClientError)
	}
	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	gin.SetMode(engineConfig.GIN_MODE)
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	user.UserController{
		UserService: &user.UserService{
			MongoContext: mongoClient,
		},
	}.RegisterController(engine)

	fmt.Println(engine.Routes())
	engine.Run(fmt.Sprintf(":%d", engineConfig.PORT))

}
