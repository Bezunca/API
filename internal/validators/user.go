package validators

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type LoginForm struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=25"`
}

type RegistrationForm struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=25"`
	Name      string `json:"name" validate:"required,min=3,max=25"`
}

type ForgotPasswordForm struct {
	Email     string `json:"email" validate:"required,email"`
}

type ResetPasswordForm struct {
	Password  string `json:"password" validate:"required,min=3,max=25"`
	Token     string `json:"token" validate:"required"`
}

type ConfirmRegistrationForm struct {
	Token     string `json:"token" validate:"required"`
}

func ValidateUserLogin(ctx echo.Context) (LoginForm, error) {

	userEmail, userPassword, ok := ctx.Request().BasicAuth()
	if !ok {
		return LoginForm{}, echo.NewHTTPError(http.StatusBadRequest, "Missing user's e-mail and password in HTTP Basic Auth")
	}

	loginForm := LoginForm{
		Email:    userEmail,
		Password: userPassword,
	}

	err := ValidateStruct(loginForm)
	if err != nil {
		return LoginForm{}, err
	}

	return loginForm, nil
}

func ValidateUserRegister(ctx echo.Context) (RegistrationForm, error) {

	registrationForm := RegistrationForm{}
	if err := ctx.Bind(&registrationForm); err != nil {
		return RegistrationForm{}, err
	}

	err := ValidateStruct(registrationForm)
	if err != nil {
		return RegistrationForm{}, err
	}

	return registrationForm, nil
}

func ValidateUserForgotPassword(ctx echo.Context) (ForgotPasswordForm, error) {

	forgotPasswordForm := ForgotPasswordForm{}
	if err := ctx.Bind(&forgotPasswordForm); err != nil {
		return ForgotPasswordForm{}, err
	}
	err := ValidateStruct(forgotPasswordForm)
	if err != nil {
		return ForgotPasswordForm{}, err
	}

	return forgotPasswordForm, nil
}

func ValidateUserResetPassword(ctx echo.Context) (ResetPasswordForm, error) {

	resetPasswordForm := ResetPasswordForm{}
	if err := ctx.Bind(&resetPasswordForm); err != nil {
		return ResetPasswordForm{}, err
	}
	err := ValidateStruct(resetPasswordForm)
	if err != nil {
		return ResetPasswordForm{}, err
	}

	return resetPasswordForm, nil
}

func ValidateUserConfirmRegistration(ctx echo.Context) (ConfirmRegistrationForm, error) {

	confirmRegistrationForm := ConfirmRegistrationForm{}
	if err := ctx.Bind(&confirmRegistrationForm); err != nil {
		return ConfirmRegistrationForm{}, err
	}
	err := ValidateStruct(confirmRegistrationForm)
	if err != nil {
		return ConfirmRegistrationForm{}, err
	}

	return confirmRegistrationForm, nil
}
