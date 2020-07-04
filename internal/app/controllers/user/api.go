package user

import (
	"bezuncapi/internal/config"
	"bezuncapi/internal/database"
	"bezuncapi/internal/models"
	"bezuncapi/internal/utils"
	"bezuncapi/internal/validators"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(ctx echo.Context) error {

	registrationForm, validationErrors := validators.ValidateUserRegister(ctx)
	if validationErrors != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": validationErrors})
	}

	_, err := database.GetUserByEmail(ctx, registrationForm.Email)
	if err == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": {"general": "Já existe uma conta cadastrada com esse email"}})
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(registrationForm.Password), bcrypt.MinCost)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": {"general": "Erro ao encriptar senha"}})
	}

	user := models.User{
		Name: registrationForm.Name,
		AuthCredentials: models.AuthCredentials{
			Email: registrationForm.Email,
			Password: string(hashPassword),
		},
	}

	inserted := PostUser(ctx, user)
	if !inserted {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": {"general": "Erro ao criar conta"}})
	}

	err = sendRegisterEmail(user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": {"general": "Erro ao enviar email de confirmação"}})
	}

	return ctx.JSON(http.StatusOK, nil)
}

func ConfirmRegistration(ctx echo.Context) error {

	confirmRegistrationForm, validationErrors := validators.ValidateUserConfirmRegistration(ctx)
	if validationErrors != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": validationErrors})
	}

	configs := config.Get()
	user, err := utils.ValidateToken(ctx, confirmRegistrationForm.Token, configs.JWTSecretEmail)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": {"general": "Token inválido"}})
	}

	updated := UpdateUserRegisterConfirmation(ctx, user.AuthCredentials.Email)
	if !updated {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": {"general": "Erro ao ativar conta"}})
	}

	token, err := utils.CreateToken(user, utils.AuthExpiration, configs.JWTSecretAuth)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": {"general": "Erro ao gerar token de autenticação"}})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"token": token})
}

func ForgotPassword(ctx echo.Context) error {

	forgotPasswordForm, validationErrors := validators.ValidateUserForgotPassword(ctx)
	if validationErrors != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": validationErrors})
	}

	user, err := database.GetUserByEmail(ctx, forgotPasswordForm.Email)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": {"general": "Não existe conta com esse email"}})
	}

	err = sendForgotPasswordEmail(user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": {"general": "Erro ao enviar email de redefinição"}})
	}

	return ctx.JSON(http.StatusOK, nil)
}

func ResetPassword(ctx echo.Context) error {

	resetPasswordForm, validationErrors := validators.ValidateUserResetPassword(ctx)
	if validationErrors != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": validationErrors})
	}

	configs := config.Get()
	user, err := utils.ValidateToken(ctx, resetPasswordForm.Token, configs.JWTSecretEmail)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": {"general": "Token inválido"}})
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(resetPasswordForm.Password), bcrypt.MinCost)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": {"general": "Erro ao encriptar senha"}})
	}

	updated := UpdateUserResetPassword(ctx, user.AuthCredentials.Email, string(hashPassword))
	if !updated {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": {"general": "Erro ao redefinir senha"}})
	}

	return ctx.JSON(http.StatusOK, nil)
}

func Login(ctx echo.Context) error {

	loginForm, validationErrors := validators.ValidateUserLogin(ctx)
	if validationErrors != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": validationErrors})
	}

	user, err := database.GetUserByEmail(ctx, loginForm.Email)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": {"general": "Emaill ou senha incorretos"}})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.AuthCredentials.Password), []byte(loginForm.Password))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": {"general": "Email ou senha incorretos"}})
	}

	if !user.AuthCredentials.Activated {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": {"general": "Conta não está ativada"}})
	}

	configs := config.Get()
	token, err := utils.CreateToken(user, utils.AuthExpiration, configs.JWTSecretAuth)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]map[string]string{"errors": {"general": "Erro ao gerar token de autenticação"}})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"token": token})
}

func Info(ctx echo.Context, user models.User) error {
	return ctx.JSON(http.StatusOK, user)
}
