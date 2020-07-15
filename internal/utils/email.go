package utils

import (
	"errors"
	"github.com/Bezunca/API/internal/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(receiverName, receiverEmail, subject, plainTextContent, htmlContent string) error {

	configs := config.Get()

	from := mail.NewEmail("Bezunca", "bezuncainvestimentos@gmail.com")
	to := mail.NewEmail(receiverName, receiverEmail)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(configs.SendGridAPIKEY)
	response, err := client.Send(message)
	if err != nil {
		return err
	}
	if response.StatusCode != 202 {
		return errors.New("falha ao enviar email de confirmação de cadastro")
	}
	return nil
}
