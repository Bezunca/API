package validators

import (
	"github.com/labstack/echo/v4"
)

type LoginForm struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=25"`
}

type RegistrationForm struct {
	Email                 string `json:"email" validate:"required,email"`
	Password              string `json:"password" validate:"required,min=3,max=25"`
	PasswordConfirmation  string `json:"password_confirmation"`
	Name                  string `json:"name" validate:"required,min=3,max=25"`
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

func ValidateUserLogin(ctx echo.Context) (LoginForm, map[string]string) {

	userEmail, userPassword, ok := ctx.Request().BasicAuth()
	if !ok {
		return LoginForm{}, map[string]string{"general": "Formulário inválido"}
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

func ValidateUserRegister(ctx echo.Context) (RegistrationForm, map[string]string) {

	registrationForm := RegistrationForm{}
	if err := ctx.Bind(&registrationForm); err != nil {
		return RegistrationForm{}, map[string]string{"general": "Formulário inválido"}
	}

	err := ValidateStruct(registrationForm)
	if err != nil {
		return RegistrationForm{}, err
	}

	if registrationForm.Password != registrationForm.PasswordConfirmation {
		return RegistrationForm{}, map[string]string{"password_confirmation": "Precisa ser igual a senha"}
	}

	return registrationForm, nil
}

func ValidateUserForgotPassword(ctx echo.Context) (ForgotPasswordForm, map[string]string) {

	forgotPasswordForm := ForgotPasswordForm{}
	if err := ctx.Bind(&forgotPasswordForm); err != nil {
		return ForgotPasswordForm{}, map[string]string{"general": "Formulário inválido"}
	}
	err := ValidateStruct(forgotPasswordForm)
	if err != nil {
		return ForgotPasswordForm{}, err
	}

	return forgotPasswordForm, nil
}

func ValidateUserResetPassword(ctx echo.Context) (ResetPasswordForm, map[string]string) {

	resetPasswordForm := ResetPasswordForm{}
	if err := ctx.Bind(&resetPasswordForm); err != nil {
		return ResetPasswordForm{}, map[string]string{"general": "Formulário inválido"}
	}
	err := ValidateStruct(resetPasswordForm)
	if err != nil {
		return ResetPasswordForm{}, err
	}

	return resetPasswordForm, nil
}

func ValidateUserConfirmRegistration(ctx echo.Context) (ConfirmRegistrationForm, map[string]string) {

	confirmRegistrationForm := ConfirmRegistrationForm{}
	if err := ctx.Bind(&confirmRegistrationForm); err != nil {
		return ConfirmRegistrationForm{}, map[string]string{"general": "Formulário inválido"}
	}
	err := ValidateStruct(confirmRegistrationForm)
	if err != nil {
		return ConfirmRegistrationForm{}, err
	}

	return confirmRegistrationForm, nil
}
