package validators

import (
	"github.com/labstack/echo/v4"
)

type CEISyncForm struct {
	User     string `json:"user" validate:"required,cpf"`
	Password string `json:"password" validate:"required"`
}

func ValidateCEISync(ctx echo.Context) (CEISyncForm, map[string]string) {

	ceiSyncForm := CEISyncForm{}
	if err := ctx.Bind(&ceiSyncForm); err != nil {
		ctx.Logger().Error(err)
		return CEISyncForm{}, map[string]string{"general": "Formulário inválido"}
	}
	err := ValidateStruct(ceiSyncForm)
	if err != nil {
		return CEISyncForm{}, err
	}

	return ceiSyncForm, nil
}
