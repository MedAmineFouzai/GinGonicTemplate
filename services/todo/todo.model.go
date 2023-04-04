package Todo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PRIORITY int32

const (
	URGENT PRIORITY = 2
	NORMAL PRIORITY = 1
	DEFER  PRIORITY = 0
)

type Todo struct {
	ID          primitive.ObjectID `bson:"_id"`
	NAME        string             `bson:"name,omitempty"`
	PRIORITY    PRIORITY           `bson:"priority,omitempty"`
	START_DATE  string             `bson:"start_date,omitempty"`
	END_DATE    string             `bson:"end_date,omitempty"`
	ASSIGNED_To string             `bson:"assigned_to,omitempty"`
	CREATED_AT  string             `bson:"createdAt,omitempty"`
	UPDATED_AT  string             `bson:"updatedAt,omitempty"`
}

func (task Todo) InsertOne(mongoContext *mongo.Client, DATABASE string, COLLECTION string) *Todo {
	_, insertOneResultError := mongoContext.Database(DATABASE).Collection(COLLECTION).InsertOne(context.TODO(), task)
	if insertOneResultError != nil {
		panic(insertOneResultError)
	}
	return &task

}

func (_ Todo) FindAll(mongoContext *mongo.Client, DATABASE string, COLLECTION string) *[]Todo {

	var tasks []Todo

	cursor, cursorError := mongoContext.Database(DATABASE).Collection(COLLECTION).Find(context.TODO(), bson.D{})
	if cursorError != nil {
		panic(cursorError)
	}

	if cursorError = cursor.All(context.TODO(), &tasks); cursorError != nil {
		panic(cursorError)
	}
	return &tasks

}

func (task Todo) FindOne(mongoContext *mongo.Client, DATABASE string, COLLECTION string) *Todo {
	singleResult := mongoContext.Database(DATABASE).Collection(COLLECTION).FindOne(context.TODO(), bson.M{
		"_id": task.ID,
	})
	if singleResultError := singleResult.Decode(task); singleResultError != nil {
		panic(singleResultError)
	}
	return &task

}

func (task Todo) UpdateOne(mongoContext *mongo.Client, DATABASE string, COLLECTION string) *Todo {
	_, updateResultError := mongoContext.Database(DATABASE).Collection(COLLECTION).UpdateOne(context.TODO(), bson.M{
		"_id": task.ID,
	}, task)
	if updateResultError != nil {
		panic(updateResultError)
	}
	return &task

}

func (task Todo) DeleteOne(mongoContext *mongo.Client, DATABASE string, COLLECTION string) *Todo {
	_, DeleteResultError := mongoContext.Database(DATABASE).Collection(COLLECTION).DeleteOne(context.TODO(), bson.M{
		"_id": task.ID,
	})
	if DeleteResultError != nil {
		panic(DeleteResultError)
	}

	return &task

}
