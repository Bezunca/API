package user

import (
	"bezuncapi/internal/config"
	"bezuncapi/internal/models"
	"bezuncapi/internal/utils"
)

func sendRegisterEmail(user models.User) error {

	configs := config.Get()
	token, err := utils.CreateToken(user, configs.JWTSecretEmail)
	if err != nil {
		return err
	}

	subject := "Confirmação de Cadastro"
	plainTextContent := "Bem Vindo!"
	htmlContent := "<a href='http://" + configs.ApplicationAddress() + "/user/confirm_registration/" + token + "'>CONFIRMAR CADASTRO</a>"

	err = utils.SendEmail(user.Name, user.AuthCredentials.Email, subject, plainTextContent, htmlContent)
	if err != nil {
		return err
	}

	return nil
}
