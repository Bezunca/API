package auth

import (
	"bezuncapi/internal/config"
	"bezuncapi/internal/models"
	"bezuncapi/internal/utils"
	"net/url"
	"time"
)

func generateDynamicLink(configs *config.Config, params map[string]string) string {

	dl := configs.DynamicLink + "?"
	for key, value := range params {
		dl += key + "=" + value + "&"
	}

	//TODO: IOS Flutter DynamicLink
	dl += "apn=" + configs.FlutterAndroidAppID

	return dl
}

func sendRegisterEmail(user models.User) error {

	configs := config.Get()
	tokenExpiration := time.Now().Add(utils.EmailExpiration).Unix()
	token, err := utils.CreateToken(user, tokenExpiration, configs.JWTSecretEmail)
	if err != nil {
		return err
	}

	subject := "Confirmação de Cadastro"
	plainTextContent := "Bem Vindo!"

	innerLink := url.QueryEscape(configs.WebURL + "confirm_registration?token=" + token)
	dynamicLink := generateDynamicLink(configs, map[string]string{
		"link": innerLink,
	})

	htmlContent := "<a href='" + dynamicLink + "'>CONFIRMAR CADASTRO</a>"

	err = utils.SendEmail(user.Name, user.AuthCredentials.Email, subject, plainTextContent, htmlContent)
	if err != nil {
		return err
	}

	return nil
}

func sendForgotPasswordEmail(user models.User) error {

	configs := config.Get()
	tokenExpiration := time.Now().Add(utils.EmailExpiration).Unix()
	token, err := utils.CreateToken(user, tokenExpiration, configs.JWTSecretEmail)
	if err != nil {
		return err
	}

	subject := "Redefinição de Senha"
	plainTextContent := "Redefinição"

	innerLink := url.QueryEscape(configs.WebURL + "reset_password?token=" + token)
	dynamicLink := generateDynamicLink(configs, map[string]string{
		"link": innerLink,
	})

	htmlContent := "<a href='" + dynamicLink + "'>REDEFINIR SENHA</a>"

	err = utils.SendEmail(user.Name, user.AuthCredentials.Email, subject, plainTextContent, htmlContent)
	if err != nil {
		return err
	}

	return nil
}
