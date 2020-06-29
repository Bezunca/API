package database

import (
	"bezuncapi/internal/config"
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const UserDatabase = "bezunca"
const UsersCollection = "users"

func GetConnection() (*mongo.Client, error) {
	configs := config.Get()
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	mongoClient, err := mongo.Connect(
		ctx, options.Client().ApplyURI(
			fmt.Sprintf(
				"mongodb://%s:%s@%s/?appname=Daily%%20Refresh%%20Job",
				configs.MongoUser,
				configs.MongoPassword,
				configs.MongoAddress(),
			),
		),
	)

	if err != nil {
		return nil, err
	}

	return mongoClient, nil
}

func FindDocuments(ctx echo.Context, database, collection string, filter bson.M, findOptions *options.FindOptions, parser func(*mongo.Cursor) (interface{}, bool)) (interface{}, error) {

	mongoClient := ctx.Get("mongoClient").(*mongo.Client)
	queryCollection := mongoClient.Database(database).Collection(collection)
	queryCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	cursor, err := queryCollection.Find(queryCtx, filter, findOptions)
	if err != nil {
		return nil, errors.New("collection not found")
	}

	data, ok := parser(cursor)
	if !ok {
		return nil, errors.New("cant decode data of collection")
	}

	return data, nil
}

func InsertDocuments(ctx echo.Context, database, collection string, documents []interface{}) bool {

	mongoClient := ctx.Get("mongoClient").(*mongo.Client)
	queryCollection := mongoClient.Database(database).Collection(collection)
	queryCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := queryCollection.InsertMany(queryCtx, documents)
	if err != nil {
		return false
	}

	return true
}

func UpdateDocuments(ctx echo.Context, database, collection string, filter bson.M, update bson.D) bool {

	mongoClient := ctx.Get("mongoClient").(*mongo.Client)
	queryCollection := mongoClient.Database(database).Collection(collection)
	queryCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := queryCollection.UpdateMany(queryCtx, filter, update)
	if err != nil {
		return false
	}

	return true
}
