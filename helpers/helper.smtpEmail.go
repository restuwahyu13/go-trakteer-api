package helpers

import (
	"encoding/json"
	"fmt"
	"net/smtp"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func SmtpEmail(to []string, template string) error {
	smtp_host := viper.GetString("SMTP_HOST")
	smtp_port, _ := strconv.Atoi(viper.GetString("SMTP_PORT"))
	smtp_username := viper.GetString("SMTP_USERNAME")
	smtp_password := viper.GetString("SMTP_PASSWORD")

	smtpAuth := smtp.PlainAuth("", smtp_username, smtp_password, smtp_host)
	smtpAddress := fmt.Sprintf("%s:%d", smtp_host, smtp_port)
	smtpFromEmail := viper.GetString("SMTP_EMAIL")

	htmlTemplate, _ := json.Marshal(template)
	smtpEmailError := smtp.SendMail(smtpAddress, smtpAuth, smtpFromEmail, to, htmlTemplate)

	if smtpEmailError != nil {
		logrus.Errorf("Sending email using SMTP error: %v", smtpEmailError)
	}

	return smtpEmailError
}
