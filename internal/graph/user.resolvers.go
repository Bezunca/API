package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
//go:generate go run github.com/99designs/gqlgen

import (
	"context"
	internalContext "github.com/Bezunca/API/internal/app/context"
	"github.com/Bezunca/API/internal/config"
	"github.com/Bezunca/API/internal/graph/generated"
	"github.com/Bezunca/API/internal/graph/model"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Bezunca/mongo_connection"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *queryResolver) User(c context.Context) (*model.User, error) {
	configs := config.Get()
	ctx := &internalContext.GraphQLContext{Context: c}

	c, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	user := new(model.User)
	result := mongo_connection.Get().
		Database(configs.Database.Name).
		Collection(configs.Database.UserCollection).
		FindOne(
			c,
		[]bson.E{
			{Key: "_id", Value: ctx.User().ID},
		},
	)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, result.Err()
	}

	if err := result.Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
