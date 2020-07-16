package wallet

import (
	"github.com/Bezunca/API/internal/app/context"
	"github.com/Bezunca/API/internal/message_broker"
	"github.com/Bezunca/API/internal/models"
	"github.com/Bezunca/API/internal/utils"
	"github.com/Bezunca/API/internal/validators"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func CEICredentials(c echo.Context) error {
	ctx := c.(*context.BezuncAPIContext)
	user := ctx.User()

	ceiSyncFrom, validationErrors := validators.ValidateCEISync(ctx)
	if validationErrors != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": validationErrors})
	}

	encryptedPassword, err := utils.RSAEncript(ceiSyncFrom.Password, utils.CEIPassword)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": {"general": "Erro ao encriptar senha"}})
	}

	ceiCredentials := models.CEI{
		User:     ceiSyncFrom.User,
		Password: encryptedPassword,
		Status: models.SyncStatus{
			StatusType: models.StatusPending,
			StatusMessage: "Sincronização em andamento",
			StatusDate: primitive.NewDateTimeFromTime(time.Now()),
		},
	}

	updated := UpdateCEI(ctx, user.AuthCredentials.Email, ceiCredentials)
	if !updated {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": {"general": "Erro ao salvar credenciais"}})
	}

	user.WalletsCredentials.Cei = &ceiCredentials

	UserScrapingInfo := models.Scraping{
		ID: user.ID,
		WalletsCredentials: user.WalletsCredentials,
	}

	err = message_broker.SendCEIScrapingRequest(ctx, UserScrapingInfo)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": {"general": "Erro ao enviar pedido de sincronização"}})
	}

	return ctx.JSON(http.StatusOK, nil)
}

func CEISync(c echo.Context) error {

	ctx := c.(*context.BezuncAPIContext)
	user := ctx.User()

	UserScrapingInfo := models.Scraping{
		ID: user.ID,
		WalletsCredentials: user.WalletsCredentials,
	}

	err := message_broker.SendCEIScrapingRequest(ctx, UserScrapingInfo)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": {"general": "Erro ao enviar pedido de sincronização"}})
	}

	return ctx.JSON(http.StatusOK, nil)
}
