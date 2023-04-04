package user

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type GENDER int32

const (
	MALE   GENDER = 0
	FEMALE GENDER = 1
	OTHER  GENDER = -1
)

type User struct {
	ID              int `bson:"_id"`
	EMAIL           string
	PASSWORD        string
	FULL_NAME       string
	AGE             int
	GENDER          GENDER
	PROFILE_PICTURE string
}

func (user User) Save(mongoContext *mongo.Client) *User {
	_, insertOneResultError := mongoContext.Database("alpha").Collection("User").InsertOne(context.Background(), user)
	if insertOneResultError != nil {
		panic(insertOneResultError)
	}
	return &user

}
