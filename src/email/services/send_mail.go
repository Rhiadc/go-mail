package services

import (
	"fmt"
	"github.com/gomail/src/domain"
	goMail "gopkg.in/mail.v2"
)

type MailSender struct {
	From      string
	Dialer    *goMail.Dialer
	EmailRepo domain.EmailRepo
}

func NewMailSender(EmailRepo domain.EmailRepo) *MailSender {
	return &MailSender{EmailRepo: EmailRepo}
}

func (ref *MailSender) Send(emailChan chan domain.Email) error {
	m := goMail.NewMessage()
	m.SetHeader("From", ref.From)
	for ec := range emailChan {
		m.SetHeader("Subject", ec.Subject)
		m.SetBody("text/html", ec.Body)
		for _, to := range ec.Emails {
			m.SetHeader("To", to)
			if err := ref.Dialer.DialAndSend(m); err != nil {
				fmt.Println(err)
				return err
			}
		}
		if err := ref.EmailRepo.SaveMail(ec); err != nil {
			return err
		}
	}

	return nil
}
