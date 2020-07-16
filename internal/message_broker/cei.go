package message_broker

import (
	"encoding/json"
	"errors"
	"github.com/Bezunca/API/internal/models"
	"github.com/labstack/echo/v4"
)

func SendCEIScrapingRequest(ctx echo.Context,scrappingRequest models.Scraping) error {

	rabbitMQ := Get()

	if scrappingRequest.WalletsCredentials.Cei != nil {
		request, err := json.Marshal(scrappingRequest)
		if err != nil {
			return err
		}

		err = rabbitMQ.Push(request)
		if err != nil {
			ctx.Logger().Error(err)
			return err
		}
	} else {
		ctx.Logger().Error("WARN: No configuration for CEI on user: " + scrappingRequest.ID.Hex())
		return errors.New("WARN: No configuration for CEI on user: " + scrappingRequest.ID.Hex())
	}

	return nil
}
