package mail

import (
	"gopkg.in/gomail.v2"
	"hkn-be/config"
)

type MailServiceInterface interface {
	SendEmail(to []string, subject, body string) error
}

type mailService struct {
	*config.Mailer
}

func NewMailServiceInterface(mailConfig *config.Mailer) MailServiceInterface {
	return &mailService{
		mailConfig,
	}
}

func (m mailService) SendEmail(to []string, subject, body string) error {
	x := gomail.NewMessage()
	x.SetHeader("From", m.Sender)
	x.SetHeader("To", to...)
	x.SetHeader("Subject", subject)
	x.SetBody("text/html", body)

	y := gomail.NewDialer(m.Host, m.Port, m.Username, m.Password)

	err := y.DialAndSend(x)
	if err != nil {
		return err
	}
	return nil
}
