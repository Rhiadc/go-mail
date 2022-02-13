package domain

type Email struct {
	Emails  []string `json:"emails"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

type EmailRepo interface {
	SaveMail(email Email) error
}
