package mailer

import "os"

type Mailer interface {
	Send(to, subject, body string) error
}

func NewMailer() Mailer {

	switch os.Getenv("MAILER_METHOD") {
	case "smtp":
		return NewSMTPMailer(os.Getenv("SMTP_HOST"),
			os.Getenv("SMTP_PORT"),
			os.Getenv("SMTP_USERNAME"),
			os.Getenv("SMTP_PASSWORD"),
			os.Getenv("SMTP_FROM"))
	case "dummy":
		return &DummyMailer{}
	default:
		return &DummyMailer{}
	}
}
