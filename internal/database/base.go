package database

import (
	"bezuncapi/internal/config"
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"github.com/Bezunca/mongo_connection"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"time"
)

const UserDatabase = "bezunca"
const UsersCollection = "users"

func GetConnection() (*mongo.Client, error) {

	configs := config.Get()
	caChainBytes, err := ioutil.ReadFile(configs.CAFilePath)
	if err != nil {
		return nil, err
	}
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(caChainBytes)
	if !ok {
		return nil, errors.New("unable to parse CA Chain file")
	}

	tlsConfig := &tls.Config{
		RootCAs: roots,
	}

	mongoClient, err := mongo_connection.New(&configs.MongoDB, tlsConfig)
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

	mongoClient := ctx.Get("mongoClient").(*mongo.Client)
	queryCollection := mongoClient.Database(database).Collection(collection)
	queryCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := queryCollection.InsertMany(queryCtx, documents)
	if err != nil {
		ctx.Logger().Error(err)
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
		ctx.Logger().Error(err)
		return false
	}

	return true
}
