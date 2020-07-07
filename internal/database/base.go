package database

import (
	"context"
	"errors"
	"time"

	"github.com/Bezunca/mongo_connection"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const UserDatabase = "bezunca"
const UsersCollection = "users"

func FindDocuments(ctx echo.Context, database, collection string, filter bson.M, findOptions *options.FindOptions, parser func(*mongo.Cursor) (interface{}, bool)) (interface{}, error) {

	mongoClient := mongo_connection.Get()
	queryCollection := mongoClient.Database(database).Collection(collection)
	queryCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	cursor, err := queryCollection.Find(queryCtx, filter, findOptions)
	if err != nil {
		ctx.Logger().Error(err)
		return nil, err
	}

	data, ok := parser(cursor)
	if !ok {
		return nil, errors.New("error parsing data")
	}

	return data, nil
}

func InsertDocuments(ctx echo.Context, database, collection string, documents []interface{}) bool {

	queryCollection := mongo_connection.Get().Database(database).Collection(collection)
	queryCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := queryCollection.InsertMany(queryCtx, documents)
	if err != nil {
		ctx.Logger().Error(err)
		return false
	}

	return true
}

func UpdateDocuments(ctx echo.Context, database, collection string, filter bson.M, update bson.D) bool {

	queryCollection := mongo_connection.Get().Database(database).Collection(collection)
	queryCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := queryCollection.UpdateMany(queryCtx, filter, update)
	if err != nil {
		ctx.Logger().Error(err)
		return false
	}

	return true
}
