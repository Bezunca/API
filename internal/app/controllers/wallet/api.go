package wallet

import (
	"bezuncapi/internal/app/context"
	"bezuncapi/internal/models"
	"bezuncapi/internal/utils"
	"bezuncapi/internal/validators"
	cei "github.com/Bezunca/ceilib/scraper"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CEISync(c echo.Context) error {
	ctx := c.(*context.BezuncAPIContext)
	user := ctx.User()

	ceiSyncFrom, validationErrors := validators.ValidateCEISync(ctx)
	if validationErrors != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": validationErrors})
	}

	ok, err := cei.Login(ceiSyncFrom.User, ceiSyncFrom.Password)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": {"general": "Erro ao logar no CEI"}})
	}

	if !ok {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": {"general": "Credenciais inválidas"}})
	}

	encryptedPassword, err := utils.RSAEncript(ceiSyncFrom.Password, utils.CEIPassword)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": {"general": "Erro ao encriptar senha"}})
	}

	ceiCredentials := models.CEI{
		User: ceiSyncFrom.User,
		Password: encryptedPassword,
	}

	updated := UpdateCEI(ctx, user.AuthCredentials.Email, ceiCredentials)
	if !updated {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": {"general": "Erro ao salvar credenciais"}})
	}

	return ctx.JSON(http.StatusOK, nil)
}
