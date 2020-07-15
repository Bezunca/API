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

	"github.com/Bezunca/mongo_connection"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *queryResolver) Trades(c context.Context) ([]*model.Trade, error) {
	configs := config.Get()
	ctx := &internalContext.GraphQLContext{Context: c}

	c, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	result, err := mongo_connection.Get().
		Database(configs.Database.Name).
		Collection("user_trades").
		Find(
			c,
		[]bson.E{
			{Key: "user_id", Value: ctx.User().ID},
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

func (r *queryResolver) Portfolio(c context.Context) ([]*model.PortfolioItem, error) {
	configs := config.Get()
	ctx := &internalContext.GraphQLContext{Context: c}

	c, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	result, err := mongo_connection.Get().
		Database(configs.Database.Name).
		Collection("user_portfolio").
		Find(
			c,
			[]bson.E{
				{Key: "user_id", Value: ctx.User().ID},
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

func (r *queryResolver) Dividends(c context.Context) ([]*model.Dividend, error){
	configs := config.Get()
	ctx := &internalContext.GraphQLContext{Context: c}

	c, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	result, err := mongo_connection.Get().
		Database(configs.Database.Name).
		Collection("user_dividends").
		Find(
			c,
			[]bson.E{
				{Key: "user_id", Value: ctx.User().ID},
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

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
