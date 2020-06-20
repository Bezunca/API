package database

import (
	"bezuncapi/internal/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

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
