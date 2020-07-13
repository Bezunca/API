package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	echoContext "bezuncapi/internal/app/context"
	"bezuncapi/internal/config"
	"bezuncapi/internal/graph/generated"
	"bezuncapi/internal/graph/model"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Bezunca/mongo_connection"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *queryResolver) User(c context.Context, name string) (*model.User, error) {
	configs := config.Get()
	ctx := c.(*echoContext.GraphQLContext)
	_user := ctx.User()

	fmt.Print(_user)

	var filters []bson.E
	if name != "" {
		filters = append(filters, bson.E{Key: "name", Value: name})
	}

	c, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	user := new(model.User)
	result := mongo_connection.Get().
		Database(configs.Database.Name).
		Collection(configs.Database.UserCollection).
		FindOne(c, filters)
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
