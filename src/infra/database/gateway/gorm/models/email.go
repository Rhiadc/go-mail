package models

import (
	"github.com/gomail/src/domain"
	"strings"
)

type Email struct {
	ID      string `gorm:"size:255"`
	Emails  string `gorm:"size:255"`
	Subject string `gorm:"size:255"`
	Body    string `gorm:"size:255"`
}

func (ref *Email) ToDomain() *domain.Email {
	emails := strings.Split(ref.Emails, ",")
	return &domain.Email{
		Emails:  emails,
		Subject: ref.Subject,
		Body:    ref.Body,
	}
}

func (ref *Email) FromDomain(email domain.Email) {
	ref.Emails = strings.Join(email.Emails, ",")
	ref.Subject = email.Subject
	ref.Body = email.Body
}
