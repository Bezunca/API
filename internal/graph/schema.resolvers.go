package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	internalContext "github.com/Bezunca/API/internal/app/context"
	"github.com/Bezunca/API/internal/config"
	"github.com/Bezunca/API/internal/graph/generated"
	"github.com/Bezunca/API/internal/graph/model"
	"github.com/Bezunca/mongo_connection"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	configs := config.Get()
	graphqlCtx := &internalContext.GraphQLContext{Context: ctx}

	c, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	result := mongo_connection.Get().
		Database(configs.Database.Name).
		Collection("users").
		FindOne(
			c,
			[]bson.E{
				{Key: "_id", Value: graphqlCtx.User().ID},
			},
		)
	if err := result.Err(); err != nil {
		return nil, err
	}

	var _user map[string]interface{}
	var user model.User
	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}
	_ = result.Decode(&_user)

	return &user, nil
}

func (r *userResolver) Dividends(ctx context.Context, obj *model.User) ([]*model.Dividend, error) {
	configs := config.Get()
	c, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	result, err := mongo_connection.Get().
		Database(configs.Database.Name).
		Collection("user_dividends").
		Find(
			c,
			[]bson.E{
				{Key: "user_id", Value: obj.ID},
			},
		)
	if err != nil {
		return nil, err
	}

	var dividends []*model.Dividend
	for result.Next(context.TODO()) {
		var dividendBody model.DividendBody
		err := result.Decode(&dividendBody)
		if err != nil {
			return nil, err
		}
		dividends = append(dividends, dividendBody.Data)
	}

	return dividends, nil
}

func (r *userResolver) Trades(ctx context.Context, obj *model.User) ([]*model.Trade, error) {
	configs := config.Get()

	c, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	result, err := mongo_connection.Get().
		Database(configs.Database.Name).
		Collection("user_trades").
		Find(
			c,
			[]bson.E{
				{Key: "user_id", Value: obj.ID},
			},
		)
	if err != nil {
		return nil, err
	}

	var trades []*model.Trade
	for result.Next(context.TODO()) {
		var tradeBody model.TradeBody
		err := result.Decode(&tradeBody)
		if err != nil {
			return nil, err
		}
		trades = append(trades, tradeBody.Data)
	}

	return trades, nil
}

func (r *userResolver) Portfolio(ctx context.Context, obj *model.User) ([]*model.PortfolioItem, error) {
	configs := config.Get()

	c, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	result, err := mongo_connection.Get().
		Database(configs.Database.Name).
		Collection("user_portfolio").
		Find(
			c,
			[]bson.E{
				{Key: "user_id", Value: obj.ID},
			},
		)
	if err != nil {
		return nil, err
	}

	var portfolio []*model.PortfolioItem
	for result.Next(context.TODO()) {
		var portfolioBody model.PortfolioBody
		err := result.Decode(&portfolioBody)
		if err != nil {
			return nil, err
		}
		portfolio = append(portfolio, portfolioBody.Data)
	}

	return portfolio, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
