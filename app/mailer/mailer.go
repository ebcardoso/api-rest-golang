package mailer

import (
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
