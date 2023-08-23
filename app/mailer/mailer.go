package mailer

import (
	"fmt"

	"github.com/ebcardoso/api-rest-golang/config"
	"github.com/go-mail/mail"
)

type MailSender struct {
	configs *config.Config
}

func NewMailSender(configs *config.Config) *MailSender {
	return &MailSender{
		configs: configs,
	}
}

func (ms *MailSender) TokenForgotPassword(email string, token string) {
	subject := "Forgot Password Token"
	body := fmt.Sprintf("<h1>Use this token to recover the password:</h1><br/><h2>%s</h2>", token)

	ms.deliverMail(email, subject, body)
}

// Delivering Email
func (ms *MailSender) deliverMail(to string, subject string, message string) {
	addr := ms.configs.Env.SMTP_ADDRESS
	port := ms.configs.Env.SMTP_PORT
	name := ms.configs.Env.SMTP_SENDER_NAME
	from := ms.configs.Env.SMTP_USERNAME
	pass := ms.configs.Env.SMTP_PASSWORD

	m := mail.NewMessage()
	m.SetAddressHeader("From", from, name)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", message)

	d := mail.NewDialer(addr, port, from, pass)
	d.DialAndSend(m)
}
