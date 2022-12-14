package helpers

import (
	"fmt"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func SmtpEmail(to []string, subject, template string) error {
	smtp_host := viper.GetString("SMTP_HOST")
	smtp_port, _ := strconv.Atoi(viper.GetString("SMTP_PORT"))
	smtp_username := viper.GetString("SMTP_USERNAME")
	smtp_password := viper.GetString("SMTP_PASSWORD")

	smtpAuth := smtp.PlainAuth("", smtp_username, smtp_password, smtp_host)
	smtpAddress := fmt.Sprintf("%s:%d", smtp_host, smtp_port)
	smtpFromEmail := viper.GetString("SMTP_EMAIL")

	smtpEmailError := smtp.SendMail(smtpAddress, smtpAuth, smtpFromEmail, to, smtpEmailMetadata(smtpFromEmail, to, subject, template))
	if smtpEmailError != nil {
		logrus.Errorf("Sending email using SMTP error: %v", smtpEmailError)
	}

	return smtpEmailError
}

func smtpEmailMetadata(from string, to []string, subject string, template string) []byte {
	mimeType := "Content-Type: text/html; \n"
	fromEmail := "From: " + from + "\n"
	toEmail := "To: " + strings.Join(to, ",") + "\n"
	subjecEmail := "Subject: " + subject + "\n"
	bodyEmail := []byte(fromEmail + toEmail + subjecEmail + mimeType + "\n" + template)
	return bodyEmail
}
